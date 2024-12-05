const std = @import("std");
const ArrayList = std.ArrayList;
const HashMap = std.StringHashMap;
const Tuple = std.meta.Tuple;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;
const Graph = @import("internal").Graph;

const inputStr = @embedFile("./in.txt");

const Manual = struct {
    rules: Graph,
    updates: std.ArrayList([]u32),

    pub fn deinit(self: *Manual) void {
        self.updates.deinit();
        self.rules.deinit();
    }
};

fn parse(allocator: std.mem.Allocator, input: []const u8) !Manual {
    var res: Manual = undefined;

    res.rules = try Graph.init(allocator, 100);
    res.updates = ArrayList([]u32).init(allocator);

    var lines = std.mem.splitScalar(u8, input, '\n');

    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }

        var splitted = std.mem.splitScalar(u8, line, '|');
        const first = try std.fmt.parseInt(u32, splitted.next().?, 10);
        const second = try std.fmt.parseInt(u32, splitted.next().?, 10);

        try res.rules.add_edge(first, second);
    }
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }
        var list = ArrayList(u32).init(allocator);
        var it = std.mem.splitScalar(u8, line, ',');
        while (it.next()) |val| {
            const num = try std.fmt.parseInt(u32, val, 10);
            try list.append(num);
        }

        try res.updates.append(try list.toOwnedSlice());
    }

    return res;
}

fn solve(allocator: std.mem.Allocator, log: Logger, manual: Manual) !u64 {
    var count: u64 = 0;
    for (manual.updates.items) |update| {
        var ordered = try allocator.alloc(u32, update.len);
        defer allocator.free(ordered);

        for (0..update.len) |currIdx| {
            const rules = manual.rules.adjacency_list[@intCast(update[currIdx])];

            var found: usize = 0;
            for (0..update.len) |idx| {
                if (currIdx == idx) continue;
                for (rules.items) |rule| {
                    if (rule == update[idx]) {
                        found += 1;
                    }
                }
            }

            const orderedIdx: usize = update.len - 1 - found;
            ordered[orderedIdx] = update[currIdx];
        }
        log.Debug("Update: {any} Ordered: {any}", .{ update, ordered });

        if (std.mem.eql(u32, update, ordered)) {
            const midx: usize = update.len / 2;
            log.Debug("Mid: {} {}", .{ ordered[midx], midx });
            count += ordered[midx];
        }
    }

    return count;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();
    var manual = try parse(allocator, inputStr);
    defer manual.deinit();

    const res = try solve(allocator, Log, manual);

    Log.Info("{}", .{res});
}
