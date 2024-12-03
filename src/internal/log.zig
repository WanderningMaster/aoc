const std = @import("std");
const Level = @import("std").log.Level;

const defaultLevel = Level.info;

const GREEN = "\x1B[32m";
const PURPLE = "\x1B[35m";
const RED = "\x1B[31m";
const ED_OFF = "\x1B[m";

pub fn NewLogger() !Log {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    const args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, args);

    var logInstance = Log{};

    if (args.len < 2) {
        return logInstance;
    }
    const maybeLevel: ?Level = std.meta.stringToEnum(Level, args[1]);
    if (maybeLevel) |level| {
        logInstance.log_level = level;
    }

    return logInstance;
}

pub const Log = struct {
    log_level: Level = defaultLevel,

    fn CheckLevel(self: Log, lvl: Level) bool {
        return @intFromEnum(self.log_level) >= @intFromEnum(lvl);
    }

    pub fn Info(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.info)) {
            std.debug.print(GREEN ++ format ++ ED_OFF ++ "\n", args);
        }
    }
    pub fn Debug(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.debug)) {
            std.debug.print(PURPLE ++ format ++ ED_OFF ++ "\n", args);
        }
    }
    pub fn DebugErr(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.debug)) {
            std.debug.print(RED ++ format ++ ED_OFF ++ "\n", args);
        }
    }
    pub fn Warn(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.warn)) {
            std.debug.print(RED ++ format ++ ED_OFF ++ "\n", args);
        }
    }
    pub fn Err(self: Log, comptime format: []const u8, args: anytype) void {
        if (self.CheckLevel(.err)) {
            std.debug.print(RED ++ format ++ ED_OFF ++ "\n", args);
        }
    }
};
