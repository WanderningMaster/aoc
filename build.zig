const std = @import("std");
const Level = @import("std").log.Level;

pub fn build(b: *std.Build) void {
    const day: usize = b.option(usize, "d", "Selected day") orelse 1;
    const year: usize = b.option(usize, "y", "Selected year") orelse 2024;
    const part: usize = b.option(usize, "p", "Selected part") orelse 1;
    const logLevel: ?[]const u8 = b.option([]const u8, "level", "Selected level");

    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const BOLD = comptime "\x1B[1m";
    const ED_OFF = comptime "\x1B[m";

    std.debug.print("{s}Year {} Day {} Part {}{s}\n", .{ BOLD, year, day, part, ED_OFF });
    std.debug.print("\n{s}Output: {s}", .{ BOLD, ED_OFF });

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

    if (logLevel) |lvl| {
        run_cmd.addArg(lvl);
    }

    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
}
