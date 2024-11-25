const std = @import("std");
const internal = @import("internal");

const data = @embedFile("./data.txt");

pub fn main() void {
    const sum = internal.math.add(1, 9);
    std.debug.print("{}\n", .{sum});
}
