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

const MapError = error{Looped};

const Map = struct {
    guard: Point,
    obstructions: std.AutoHashMap(Point, u1),
    currVec: Point,
    distinctLocations: std.AutoHashMap(Point, u1),
    allocator: std.mem.Allocator,
    boundaries: Point,
    log: Logger,
    loopGuard: std.AutoHashMap(Point, Point),

    pub fn deinit(self: *Map) void {
        self.distinctLocations.deinit();
        self.obstructions.deinit();
        self.loopGuard.deinit();
    }

    pub fn visit(self: *Map, x: i32, y: i32) !bool {
        const key = Point{ .x = x, .y = y };

        if (self.loopGuard.contains(key)) {
            const vec = self.loopGuard.get(key).?;
            if (vec.x == self.currVec.x and vec.y == self.currVec.y) {
                self.log.DebugErr("Loop detected", .{});
                return true;
            }
        } else {
            try self.loopGuard.put(key, self.currVec);
        }

        if (!self.distinctLocations.contains(key)) {
            try self.distinctLocations.put(key, 0);
        }
        self.guard = key;
        return false;
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

        const loop = try self.visit(nextX, nextY);
        if (loop) {
            return MapError.Looped;
        }
        return true;
    }

    pub fn init(allocator: std.mem.Allocator, log: Logger) Map {
        var map: Map = undefined;
        map.distinctLocations = std.AutoHashMap(Point, u1).init(allocator);
        map.obstructions = std.AutoHashMap(Point, u1).init(allocator);
        map.allocator = allocator;
        map.boundaries = Point{ .x = 0, .y = 0 };
        map.loopGuard = std.AutoHashMap(Point, Point).init(allocator);
        map.log = log;

        return map;
    }

    pub fn clone(allocator: std.mem.Allocator, log: Logger, map: *Map) !Map {
        var newMap = Map.init(allocator, log);

        newMap.obstructions = try map.obstructions.clone();
        newMap.distinctLocations = try map.distinctLocations.clone();
        newMap.guard = map.guard;
        newMap.currVec = map.currVec;
        newMap.boundaries = map.boundaries;

        return newMap;
    }
};

fn parse(allocator: std.mem.Allocator, log: Logger, input: []const u8) !Map {
    var map = Map.init(allocator, log);
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
                    _ = try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Up [{}][{}]", .{ x, y });
                },
                '>' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = 1, .y = 0 };
                    _ = try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Right [{}][{}]", .{ x, y });
                },
                'v' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = 0, .y = 1 };
                    _ = try map.visit(map.guard.x, map.guard.y);
                    log.Debug("Guard Down [{}][{}]", .{ x, y });
                },
                '<' => {
                    map.guard = Point{
                        .x = @intCast(x),
                        .y = @intCast(y),
                    };
                    map.currVec = Point{ .x = -1, .y = 0 };
                    _ = try map.visit(map.guard.x, map.guard.y);
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

fn solve(map: *Map, allocator: std.mem.Allocator, log: Logger) !u32 {
    var hitLoop: u32 = 0;

    var mapCopy = try Map.clone(allocator, log, map);
    while (try mapCopy.move()) {}

    var it = mapCopy.distinctLocations.keyIterator();
    while (it.next()) |val| {
        const p = Point{ .x = val.x, .y = val.y };

        var modified = try Map.clone(allocator, log, map);

        if (!modified.obstructions.contains(p)) {
            try modified.obstructions.put(p, 0);
        }

        while (true) {
            const stop = modified.move() catch {
                hitLoop += 1;
                break;
            };
            if (!stop) break;
        }

        modified.deinit();
    }

    // for (0..@intCast(map.boundaries.x + 1)) |x| {
    //     for (0..@intCast(map.boundaries.y + 1)) |y| {
    //         const p = Point{ .x = @intCast(x), .y = @intCast(y) };
    //
    //         var newMap = try Map.clone(allocator, log, map);
    //
    //         if (!newMap.obstructions.contains(p)) {
    //             try newMap.obstructions.put(p, 0);
    //         }
    //
    //         while (true) {
    //             const stop = newMap.move() catch {
    //                 // log.Err("Tried: {any}", .{p});
    //                 hitLoop += 1;
    //                 break;
    //             };
    //             if (!stop) break;
    //         }
    //
    //         newMap.deinit();
    //     }
    // }

    return hitLoop;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    var map = try parse(allocator, Log, inputStr);
    defer map.deinit();

    const res = try solve(&map, allocator, Log);
    Log.Info("{}", .{res});
}
