pub const math = struct {
    pub const add = @import("add.zig").add;
};
pub const NewLogger = @import("log.zig").NewLogger;
pub const Logger = @import("log.zig").Log;
