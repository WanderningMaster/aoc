const std = @import("std");
const ArrayList = std.ArrayList;
const HashMap = std.StringHashMap;
const Tuple = std.meta.Tuple;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;

const inputStr = @embedFile("./in.txt");

fn parse(allocator: std.mem.Allocator, input: []const u8) ![][]u8 {
    var lines = std.mem.splitScalar(u8, input, '\n');
    var grid = ArrayList([]u8).init(allocator);

    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }
        var row = ArrayList(u8).init(allocator);
        try row.appendSlice(line);
        try grid.append(try row.toOwnedSlice());
    }
    return try grid.toOwnedSlice();
}

fn searchAtPos(allocator: std.mem.Allocator, grid: [][]u8, startX: usize, startY: usize) !u32 {
    var count: u32 = 0;

    const x: i64 = @as(i64, @intCast(startX));
    const y: i64 = @as(i64, @intCast(startY));
    if (x - 1 < 0 or x + 1 >= grid.len or y - 1 < 0 or y + 1 >= grid[0].len) {
        return count;
    }

    var dRight = ArrayList(u8).init(allocator);
    defer dRight.deinit();

    try dRight.append(grid[startX - 1][startY - 1]);
    try dRight.append(grid[startX][startY]);
    try dRight.append(grid[startX + 1][startY + 1]);

    var dLeft = ArrayList(u8).init(allocator);
    defer dLeft.deinit();

    try dLeft.append(grid[startX - 1][startY + 1]);
    try dLeft.append(grid[startX][startY]);
    try dLeft.append(grid[startX + 1][startY - 1]);

    if ((std.mem.eql(u8, dLeft.items, "MAS") or std.mem.eql(u8, dLeft.items, "SAM")) and
        (std.mem.eql(u8, dRight.items, "MAS") or std.mem.eql(u8, dRight.items, "SAM")))
    {
        count += 1;
    }

    return count;
}

fn solve(allocator: std.mem.Allocator, grid: [][]u8, log: Logger) !u32 {
    var sum: u32 = 0;
    for (0..grid.len) |i| {
        for (0..grid[i].len) |j| {
            if (grid[i][j] == 'A') {
                log.Debug("Found possible cross on [{}][{}]", .{ i, j });
                sum += try searchAtPos(allocator, grid, i, j);
            }
        }
    }

    return sum;
}

const Level = @import("std").log.Level;
pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    const grid = try parse(allocator, inputStr);

    const res = try solve(allocator, grid, Log);
    Log.Info("{}", .{res});
}
