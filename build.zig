const std = @import("std");

pub fn build(b: *std.Build) void {
    const day: usize = b.option(usize, "d", "Selected day") orelse @panic("Should select day");
    const year: usize = b.option(usize, "y", "Selected year") orelse @panic("Should select day");
    const part: usize = b.option(usize, "p", "Selected part") orelse @panic("Should select part");

    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const BOLD = comptime "\x1B[1m";
    const GREEN = comptime "\x1B[32m";
    const ED_OFF = comptime "\x1B[m";

    std.debug.print("{s}Year {} Day {} Part {}{s}\n", .{ BOLD, year, day, part, ED_OFF });
    std.debug.print("\n{s}Output: {s}", .{ BOLD, ED_OFF });

    std.debug.print("{s}\n", .{GREEN});

    const source_path = b.fmt("src/{}/day{}/part{}/main.zig", .{ year, day, part });

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
}
