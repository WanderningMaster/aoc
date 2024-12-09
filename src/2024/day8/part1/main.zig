const std = @import("std");
const ArrayList = std.ArrayList;
const HashMap = std.StringHashMap;
const Tuple = std.meta.Tuple;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;
const GraphGeneric = @import("internal").GraphGeneric;

const inputStr = @embedFile("./in.txt");

const Point = struct {
    x: i32,
    y: i32,

    fn add(self: Point, p2: Point) Point {
        return Point{
            .x = self.x + p2.x,
            .y = self.y + p2.y,
        };
    }

    fn sub(self: Point, p2: Point) Point {
        return Point{
            .x = self.x - p2.x,
            .y = self.y - p2.y,
        };
    }
};
const Graph = GraphGeneric(Point);

const Map = struct {
    antennas: Graph,
    boundaries: Point,
    log: Logger,
    antinodes: std.AutoHashMap(Point, u1),

    pub fn push_antinode(self: *Map, p: Point) !void {
        self.log.Debug("Adding antinode: {any}", .{p});
        if (self.antinodes.contains(p)) {
            self.log.Debug("Antinode already exists", .{});
            return;
        }

        if (p.x > self.boundaries.x or p.x < 0 or p.y > self.boundaries.y or p.y < 0) {
            self.log.Debug("Out of boundary", .{});
            return;
        }

        try self.antinodes.put(p, 0);
        self.log.Debug("Added", .{});
    }

    pub fn deinit(self: *Map) void {
        self.antennas.deinit();
        self.antinodes.deinit();
    }
};

fn parse(allocator: std.mem.Allocator, log: Logger, input: []const u8) !Map {
    var map: Map = undefined;
    map.antennas = try Graph.init(allocator, 123);
    map.boundaries = Point{ .x = 0, .y = 0 };
    map.antinodes = std.AutoHashMap(Point, u1).init(allocator);

    var lines = std.mem.splitScalar(u8, input, '\n');

    var y: usize = 0;

    while (lines.next()) |line| : (y += 1) {
        if (line.len == 0) continue;

        map.boundaries.x = @as(i32, @intCast(line.len)) - 1;
        map.boundaries.y += 1;

        var listLine = try ArrayList(u8).initCapacity(allocator, line.len);
        for (line, 0..) |ch, x| {
            switch (ch) {
                '0'...'9', 'a'...'z', 'A'...'Z' => {
                    try map.antennas.add_edge(@intCast(ch), Point{ .x = @intCast(x), .y = @intCast(y) });
                    log.Debug("Antena '{c}' [{}][{}]", .{ ch, x, y });
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
    while (map.antennas.next()) |antennas| {
        for (0..antennas.len) |i| {
            for (i + 1..antennas.len) |j| {
                const p1 = antennas[i];
                const p2 = antennas[j];

                map.log.Debug("Pair: ({}, {}) <-> ({}, {})", .{ p1.x, p1.y, p2.x, p2.y });

                const distance = p1.sub(p2);
                try map.push_antinode(p1.add(distance));
                try map.push_antinode(p2.sub(distance));
            }
        }
    }

    return map.antinodes.count();
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    var map = try parse(allocator, Log, inputStr);
    defer map.deinit();

    const res = try solve(&map);
    Log.Info("{}", .{res});
}

// -3 +1
// 8,1 p1
// 5,2 p2
//
// 11, 0 a1
// 2,3 a2
