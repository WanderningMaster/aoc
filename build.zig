const std = @import("std");

pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});
    const allocator = std.heap.page_allocator;

    const day: usize = b.option(usize, "d", "Select day") orelse @panic("Day should be specified");
    const year: usize = b.option(usize, "y", "Select year") orelse @panic("Year should be specified");

    std.debug.print("Day {} Year {}\n", .{ day, year });

    const source_path = try std.fmt.allocPrint(allocator, "{}/day{}/main.zig", .{ year, day });
    defer allocator.free(source_path);

    const exe = b.addExecutable(.{
        .name = "aoc",
        .root_source_file = b.path(source_path),
        .target = target,
        .optimize = optimize,
    });

    b.installArtifact(exe);

    const run_cmd = b.addRunArtifact(exe);

    run_cmd.step.dependOn(b.getInstallStep());

    std.debug.print("\nOutput:\n", .{});
    // This creates a build step. It will be visible in the `zig build --help` menu,
    // and can be selected like this: `zig build run`
    // This will evaluate the `run` step rather than the default, which is "install".
    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
}
