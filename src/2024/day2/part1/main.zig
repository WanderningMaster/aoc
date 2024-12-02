const std = @import("std");
const ArrayList = std.ArrayList;

const inputStr = @embedFile("./in.txt");

const Report = struct {
    data: []i32,

    pub fn SafetyCheck(self: *const Report) bool {
        var diffSign: i8 = 0;
        for (0..self.data.len - 1) |idx| {
            const diff: i32 = self.data[idx] - self.data[idx + 1];
            if (diffSign == 0) {
                diffSign = if (diff < 0) -1 else 1;
            }

            const sameSign = (diff ^ diffSign) >= 0;
            // std.debug.print("{} {}\n", .{ sameSign, diff });
            if (sameSign and @abs(diff) >= 1 and @abs(diff) <= 3) {
                continue;
            } else {
                return false;
            }
        }

        return true;
    }
};

pub fn parse(allocator: std.mem.Allocator, input: []const u8) ![]Report {
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
        const report = Report{ .data = data };
        try reports.append(report);
    }
    return try reports.toOwnedSlice();
}

pub fn solve(reports: []Report) !u32 {
    var count: u32 = 0;
    for (reports) |report| {
        count += if (report.SafetyCheck()) 1 else 0;
        // std.debug.print("{any} {}\n", .{ report.data, report.SafetyCheck() });
    }

    return count;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const reports = try parse(allocator, inputStr);
    const res = try solve(reports);

    std.debug.print("{}\n", .{res});
}
