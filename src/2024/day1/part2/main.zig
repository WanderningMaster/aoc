const std = @import("std");
const ArrayList = std.ArrayList;

const inputStr = @embedFile("./in.txt");

pub fn parse(allocator: std.mem.Allocator, input: []const u8) ![2][]i32 {
    var list1 = ArrayList(i32).init(allocator);
    var list2 = ArrayList(i32).init(allocator);
    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |line| {
        if (line.len == 0) {
            break;
        }
        var numsStr = std.mem.splitSequence(u8, line, "   ");
        const num1 = try std.fmt.parseInt(i32, numsStr.next().?, 10);
        const num2 = try std.fmt.parseInt(i32, numsStr.next().?, 10);
        try list1.append(num1);
        try list2.append(num2);
    }

    const arr1 = try list1.toOwnedSlice();
    const arr2 = try list2.toOwnedSlice();
    const lists: [2][]i32 = .{ arr1, arr2 };

    return lists;
}

pub fn occurences(list: []i32, target: i32) u64 {
    var count: u64 = 0;
    for (list) |x| {
        if (target == x) {
            count += 1;
        }
    }

    return count;
}

pub fn solve(lists: [2][]i32) !u64 {
    const list1 = lists[0];
    const list2 = lists[1];

    var distance: u64 = 0;
    for (0..list1.len) |idx| {
        const num: u64 = @intCast(list1[idx]);
        distance += num * occurences(list2, list1[idx]);
    }

    return distance;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const lists = try parse(allocator, inputStr);
    const res = try solve(lists);

    std.debug.print("{}\n", .{res});
}
