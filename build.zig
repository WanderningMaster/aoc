const std = @import("std");

pub fn build(b: *std.Build) void {
    const day: usize = b.option(usize, "d", "Selected day") orelse @panic("Should select day");
    const year: usize = b.option(usize, "y", "Selected year") orelse @panic("Should select day");

    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    std.debug.print("Year {} Day {}\n", .{ year, day });

    const source_path = b.fmt("src/{}/day{}/main.zig", .{ year, day });

    const exe = b.addExecutable(.{
        .name = "zig-aoc",
        .root_source_file = b.path(source_path),
        .target = target,
        .optimize = optimize,
    });
    const internal = b.addModule("internal", .{ .root_source_file = b.path("./src/internal/internal.zig") });
    exe.root_module.addImport("internal", internal);

    b.installArtifact(exe);
    const run_cmd = b.addRunArtifact(exe);

    run_cmd.step.dependOn(b.getInstallStep());

    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);

    std.debug.print("\nOutput:\n", .{});
}
