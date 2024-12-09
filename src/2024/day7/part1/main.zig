const std = @import("std");
const ArrayList = std.ArrayList;
const HashMap = std.StringHashMap;
const Tuple = std.meta.Tuple;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;

const inputStr = @embedFile("./in.txt");

const Op = enum(u64) {
    MUL = 0,
    ADD = 1,
    CONCAT = 2,
};

fn concatNumbers(n1: u64, n2: u64) u64 {
    const len2 = std.math.log10(n2) + 1;
    return n1 * std.math.pow(u64, 10, len2) + n2;
}

const Eq = struct {
    sum: u64,
    nums: []u64,
    log: Logger,

    fn generateOps(self: Eq, allocator: std.mem.Allocator, n: usize) ![][]Op {
        self.log.Debug("Generating for n: {}\n", .{n});
        var results = ArrayList([]Op).init(allocator);
        const totalCombinations = std.math.pow(u64, 3, @intCast(n));

        for (0..totalCombinations) |i| {
            var combination = try allocator.alloc(Op, n);
            var num = i;

            for (0..n) |idx| {
                switch (num % 3) {
                    0 => combination[idx] = Op.MUL,
                    1 => combination[idx] = Op.ADD,
                    2 => combination[idx] = Op.CONCAT,
                    else => unreachable,
                }
                num /= 3;
            }
            self.log.Debug("{any}", .{combination});

            try results.append(combination);
        }
        self.log.Debug("", .{});

        return try results.toOwnedSlice();
    }

    fn calc(self: Eq, ops: []Op) bool {
        var res: ?u64 = null;
        for (ops, 0..) |op, idx| {
            if (res == null) {
                res = self.nums[idx];
            }

            switch (op) {
                .MUL => res.? *= self.nums[idx + 1],
                .ADD => res.? += self.nums[idx + 1],
                .CONCAT => res = concatNumbers(res.?, self.nums[idx + 1]),
            }
        }

        return self.sum == res.?;
    }
};

fn parseLine(allocator: std.mem.Allocator, log: Logger, line: []const u8) !Eq {
    var it = std.mem.splitSequence(u8, line, ": ");

    const sum = try std.fmt.parseInt(u64, it.next().?, 10);
    const second = it.next().?;

    var numsIt = std.mem.splitScalar(u8, second, ' ');
    var list = ArrayList(u64).init(allocator);
    while (numsIt.next()) |numStr| {
        const num = try std.fmt.parseInt(u64, numStr, 10);
        try list.append(num);
    }

    return Eq{
        .sum = sum,
        .nums = try list.toOwnedSlice(),
        .log = log,
    };
}

fn parse(allocator: std.mem.Allocator, log: Logger, input: []const u8) ![]Eq {
    var lines = std.mem.splitScalar(u8, input, '\n');

    var list = ArrayList(Eq).init(allocator);
    while (lines.next()) |line| {
        if (line.len == 0) continue;

        const eq = try parseLine(allocator, log, line);
        try list.append(eq);
    }

    return try list.toOwnedSlice();
}

fn solve(allocator: std.mem.Allocator, eqs: []Eq) !u129 {
    var sum: u128 = 0;

    for (eqs) |eq| {
        const combinations = try eq.generateOps(allocator, eq.nums.len - 1);
        defer {
            for (combinations) |combination| {
                allocator.free(combination);
            }
        }
        for (combinations) |combination| {
            if (eq.calc(combination)) {
                eq.log.Debug("Found {any}, {}", .{ combination, eq.sum });
                sum += @intCast(@as(u128, eq.sum));

                break;
            }
        }
    }

    return sum;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    const eqs = try parse(allocator, Log, inputStr);

    const res = try solve(allocator, eqs);
    Log.Info("{}", .{res});
}
