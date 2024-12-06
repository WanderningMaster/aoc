const std = @import("std");
const ArrayList = std.ArrayList;
const HashMap = std.StringHashMap;
const Tuple = std.meta.Tuple;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;

const inputStr = @embedFile("./in.txt");

const Point = struct {
    x: i32,
    y: i32,
};

const Map = struct {
    guard: Point,
    obstructions: std.AutoHashMap(Point, u1),
    currVec: Point,
    distinctLocations: std.AutoHashMap(Point, u1),
    allocator: std.mem.Allocator,
    boundaries: Point,
    log: Logger,

    pub fn deinit(self: *Map) void {
        self.distinctLocations.deinit();
        self.obstructions.deinit();
    }

    pub fn visit(self: *Map, x: i32, y: i32) !void {
        const key = Point{ .x = x, .y = y };
        if (!self.distinctLocations.contains(key)) {
            try self.distinctLocations.put(key, 0);
        }
        self.guard = key;
    }

    pub fn hit(self: *Map, x: i32, y: i32) !bool {
        const key = Point{ .x = x, .y = y };

        return self.obstructions.contains(key);
    }

    pub fn move(self: *Map) !bool {
        const nextX: i32 = self.guard.x + self.currVec.x;
        const nextY: i32 = self.guard.y + self.currVec.y;

        if (nextX > self.boundaries.x or nextX < 0 or nextY > self.boundaries.y or nextY < 0) {
            self.log.Debug("Out of boundary", .{});
            return false;
        }

        if (try self.hit(nextX, nextY)) {
            self.log.Debug("Hit obstruction", .{});
            self.currVec = Point{ .x = self.currVec.y * -1, .y = self.currVec.x };
            return true;
        }

        self.log.Debug("Moving to ({}, {})", .{ nextX, nextY });
        try self.visit(nextX, nextY);
        return true;
    }
};

fn parse(allocator: std.mem.Allocator, log: Logger, input: []const u8) !Map {
    var map: Map = undefined;
    map.distinctLocations = std.AutoHashMap(Point, u1).init(allocator);
    map.obstructions = std.AutoHashMap(Point, u1).init(allocator);
    map.allocator = allocator;
    map.boundaries = Point{ .x = 0, .y = 0 };

    var lines = std.mem.splitScalar(u8, input, '\n');

    var y: usize = 0;

    while (lines.next()) |line| : (y += 1) {
        if (line.len == 0) continue;

        map.boundaries.x = @as(i32, @intCast(line.len)) - 1;
        map.boundaries.y += 1;

        var listLine = try ArrayList(u8).initCapacity(allocator, line.len);
        for (line, 0..) |ch, x| {
            switch (ch) {
                '#' => {
                    try map.obstructions.put(Point{ .x = @intCast(x), .y = @intCast(y) }, 0);
                    log.Debug("Obstruction [{}][{}]", .{ x, y });
                },
                '^' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = 0, .y = -1 };
                    try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Up [{}][{}]", .{ x, y });
                },
                '>' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = 1, .y = 0 };
                    try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Right [{}][{}]", .{ x, y });
                },
                'v' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = 0, .y = 1 };
                    try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Down [{}][{}]", .{ x, y });
                },
                '<' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = -1, .y = 0 };
                    try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Left [{}][{}]", .{ x, y });
                },
                else => {},
            }
            try listLine.append(ch);
        }
    }
    map.boundaries.y -= 1;
    map.log = log;

    return map;
}

fn solve(map: *Map) !u32 {
    while (try map.move()) {}

    return map.distinctLocations.count();
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    var map = try parse(allocator, Log, inputStr);
    defer map.deinit();
    Log.Info("{any}", .{130 * 130 - map.obstructions.count()});

    const res = try solve(&map);
    Log.Info("{}", .{res});
}
