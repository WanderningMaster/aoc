const std = @import("std");
const ArrayList = std.ArrayList;
const Tuple = std.meta.Tuple;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;

const inputStr = @embedFile("./in.txt");

const TupleDef = Tuple(&.{ bool, []usize });
const Report = struct {
    data: []i32,
    Log: Logger,

    pub fn SafetyCheck(self: *const Report, allocator: std.mem.Allocator) !TupleDef {
        var diffSign: i8 = 0;
        for (0..self.data.len - 1) |idx| {
            const diff: i32 = self.data[idx] - self.data[idx + 1];
            if (diffSign == 0) {
                diffSign = if (diff < 0) -1 else 1;
            }

            const sameSign = (diff ^ diffSign) >= 0;
            if (sameSign and @abs(diff) >= 1 and @abs(diff) <= 3) {
                continue;
            } else {
                var cList = ArrayList(usize).init(allocator);
                defer cList.deinit();

                try cList.append(idx);
                try cList.append(idx + 1);
                if (idx + 2 < self.data.len - 1) {
                    try cList.append(idx + 2);
                }

                const maybeNegative: i32 = @intCast(idx);
                if (maybeNegative - 1 >= 0) {
                    try cList.append(idx - 1);
                }
                const considerations = try cList.toOwnedSlice();

                return .{ false, considerations };
            }
        }

        const emptySlice: []usize = undefined;
        return .{ true, emptySlice };
    }

    fn consider(self: *const Report, allocator: std.mem.Allocator, pos: usize) !bool {
        var newList = try ArrayList(i32).fromOwnedSlice(allocator, self.data).clone();
        _ = newList.orderedRemove(pos);

        const newData = try newList.toOwnedSlice();

        const modifiedReport = Report{ .data = newData, .Log = self.Log };

        const res = try modifiedReport.SafetyCheck(allocator);
        if (res[0]) {
            return true;
        }
        return false;
    }

    pub fn SafetyCheckV2(self: *const Report, allocator: std.mem.Allocator) !bool {
        const res = try self.SafetyCheck(allocator);
        if (res[0]) {
            return true;
        }
        for (res[1]) |pos| {
            self.Log.Debug("Consider for pos {}", .{pos});
            const passed = try self.consider(allocator, pos);
            if (passed) {
                return true;
            }
        }

        return false;
    }
};

pub fn parse(allocator: std.mem.Allocator, input: []const u8, Log: Logger) ![]Report {
    var reports = ArrayList(Report).init(allocator);
    defer reports.deinit();

    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |line| {
        if (line.len == 0) {
            continue;
        }

        var list = ArrayList(i32).init(allocator);
        defer list.deinit();

        var lvls = std.mem.splitScalar(u8, line, ' ');
        while (lvls.next()) |lvl| {
            const num = try std.fmt.parseInt(i32, lvl, 10);
            try list.append(num);
        }
        const data = try list.toOwnedSlice();
        const report = Report{ .data = data, .Log = Log };
        try reports.append(report);
    }
    return try reports.toOwnedSlice();
}

pub fn solve(reports: []Report, allocator: std.mem.Allocator, Log: Logger) !u32 {
    var count: u32 = 0;
    for (reports) |report| {
        Log.Debug("Checking for {any}", .{report.data});
        const passed = try report.SafetyCheckV2(allocator);
        if (!passed) {
            Log.Debug("Failed for {any}", .{report.data});
        }

        count += if (passed) 1 else 0;
    }

    return count;
}

pub fn main() !void {
    const Log = try NewLogger();
    const allocator = std.heap.page_allocator;
    const reports = try parse(allocator, inputStr, Log);
    const res = try solve(reports, allocator, Log);

    Log.Info("{}", .{res});
}
