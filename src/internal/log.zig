const std = @import("std");
const Level = @import("std").log.Level;

pub fn NewLogger() !Log {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    const args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, args);
    if (args.len < 2) {
        return Log{};
    }

    var level: Level = undefined;
    if (std.mem.eql(u8, args[1], "debug")) {
        level = .debug;
    } else if (std.mem.eql(u8, args[1], "warn")) {
        level = .warn;
    } else if (std.mem.eql(u8, args[1], "err")) {
        level = .err;
    } else {
        level = .info;
    }

    return Log{ .log_level = level };
}

pub const Log = struct {
    log_level: Level = .info,

    fn CheckLevel(self: Log, lvl: Level) bool {
        return @intFromEnum(self.log_level) >= @intFromEnum(lvl);
    }

    pub fn Info(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.info)) {
            std.debug.print(format ++ "\n", args);
        }
    }
    pub fn Debug(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.debug)) {
            std.debug.print(format ++ "\n", args);
        }
    }
    pub fn Warn(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.warn)) {
            std.debug.print(format ++ "\n", args);
        }
    }
    pub fn Err(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.err)) {
            std.debug.print(format ++ "\n", args);
        }
    }
};
