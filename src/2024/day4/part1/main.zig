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

fn searchAtPos(allocator: std.mem.Allocator, log: Logger, grid: [][]u8, startX: usize, startY: usize) !u32 {
    var count: u32 = 0;

    const directions = [_]Tuple(&.{ i8, i8 }){
        .{ 0, 1 }, // 0 degrees: right
        .{ -1, 1 }, // 45 degrees: up-right
        .{ -1, 0 }, // 90 degrees: up
        .{ -1, -1 }, // 135 degrees: up-left
        .{ 0, -1 }, // 180 degrees: left
        .{ 1, -1 }, // 225 degrees: down-left
        .{ 1, 0 }, // 270 degrees: down
        .{ 1, 1 }, // 315 degrees: down-right}
    };

    for (directions) |direction| {
        var list = ArrayList(u8).init(allocator);
        defer list.deinit();

        const dx = direction[0];
        const dy = direction[1];

        for (0..4) |i| {
            const x: i64 = @as(i64, @intCast(startX)) + @as(i64, @intCast(dx)) * @as(i64, @intCast(i));
            const y: i64 = @as(i64, @intCast(startY)) + @as(i64, @intCast(dy)) * @as(i64, @intCast(i));
            log.Debug("[{}][{}], len: {}", .{ x, y, grid.len });
            if (x >= 0 and x < grid.len and y >= 0 and y < grid[0].len) {
                try list.append(grid[@intCast(x)][@intCast(y)]);
            } else {
                break;
            }
        }
        log.Debug("{any}", .{list.items});
        if (std.mem.eql(u8, "XMAS", list.items) or std.mem.eql(u8, "SAMX", list.items)) {
            log.Debug("Match", .{});
            count += 1;
        }
    }

    return count;
}

fn solve(allocator: std.mem.Allocator, grid: [][]u8, log: Logger) !u32 {
    var sum: u32 = 0;
    for (0..grid.len) |i| {
        for (0..grid[i].len) |j| {
            if (grid[i][j] == 'X' or grid[i][j] == 'S') {
                log.Debug("Found possible start of sequence on [{}][{}]", .{ i, j });
                sum += try searchAtPos(allocator, log, grid, i, j);
            }
        }
    }

    return sum / 2;
}

const Level = @import("std").log.Level;
pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    const grid = try parse(allocator, inputStr);

    const res = try solve(allocator, grid, Log);
    Log.Info("{}", .{res});
}
