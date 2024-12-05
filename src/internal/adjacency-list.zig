const std = @import("std");

pub const Graph = struct {
    adjacency_list: []std.ArrayList(u32), // Each vertex has a list of adjacent vertices
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator, vertex_count: usize) !Graph {
        const adjacency_list = try allocator.alloc(std.ArrayList(u32), vertex_count);
        for (adjacency_list) |*list| {
            list.* = std.ArrayList(u32).init(allocator);
        }
        return Graph{
            .adjacency_list = adjacency_list,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Graph) void {
        for (self.adjacency_list) |*list| {
            list.deinit();
        }
        self.allocator.free(self.adjacency_list);
    }

    pub fn add_edge(self: *Graph, from: usize, to: u32) !void {
        try self.adjacency_list[from].append(to);
    }

    pub fn display(self: *const Graph) void {
        for (self.adjacency_list, 0..) |list, i| {
            if (list.items.len == 0) {
                continue;
            }
            std.debug.print("{} -> ", .{i});
            for (list.items) |vertex| {
                std.debug.print("{} ", .{vertex});
            }
            std.debug.print("\n", .{});
        }
    }
};

pub const GraphSet = struct {
    adjacency_set: []std.AutoHashMap(u32, u1), // Each vertex has a list of adjacent vertices
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator, vertex_count: usize) !GraphSet {
        const adjacency_set = try allocator.alloc(std.AutoHashMap(u32, u1), vertex_count);
        for (adjacency_set) |*set| {
            set.* = std.AutoHashMap(u32, u1).init(allocator);
        }
        return GraphSet{
            .adjacency_set = adjacency_set,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *GraphSet) void {
        for (self.adjacency_set) |*set| {
            set.deinit();
        }
        self.allocator.free(self.adjacency_set);
    }

    pub fn add_edge(self: *GraphSet, from: usize, to: u32) !void {
        if (!self.adjacency_set[from].contains(to)) {
            try self.adjacency_set[from].put(to, 0);
        }
    }
};
