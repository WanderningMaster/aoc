pub const math = struct {
    pub const add = @import("add.zig").add;
};
pub const NewLogger = @import("log.zig").NewLogger;
pub const Logger = @import("log.zig").Log;
pub const Graph = @import("adjacency-list.zig").Graph;
pub const GraphSet = @import("adjacency-list.zig").GraphSet;
pub const GraphGeneric = @import("adjacency-list.zig").GraphGeneric;
