const std = @import("std");
const internal = @import("internal");

pub fn main() void {
    const sum = internal.add(1, 2);
    std.debug.print("Hello, {}", .{sum});
}
