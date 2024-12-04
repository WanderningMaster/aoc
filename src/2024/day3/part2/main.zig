const std = @import("std");
const ArrayList = std.ArrayList;
const NewLogger = @import("internal").NewLogger;
const Logger = @import("internal").Logger;
const Tuple = std.meta.Tuple;

const inputStr = @embedFile("./in.txt");

const TokenTyp = enum { MUL, NUMBER, LBRACE, RBRACE, SPACE, NEWLINE, COMMA, ILLEGAL, DO, DONT };

const Token = struct {
    type: TokenTyp,
    value: []const u8,

    pub fn String(self: Token, allocator: std.mem.Allocator) ![]u8 {
        if (self.type == .NEWLINE) {
            return try std.fmt.allocPrint(allocator, "Token(.type={}, .value='\\n')", .{self.type});
        }
        return try std.fmt.allocPrint(allocator, "Token(.type={}, .value='{s}')", .{ self.type, self.value });
    }
};

fn parseNumber(input: []const u8, pos: usize) Tuple(&.{ Token, usize }) {
    const startPos = pos;
    var newPos = pos;
    while (input[newPos] <= '9' and input[newPos] >= '0') {
        if (pos >= input.len) {
            break;
        }
        newPos += 1;
    }

    return .{ Token{ .type = .NUMBER, .value = input[startPos..newPos] }, newPos };
}

fn tokenize(allocator: std.mem.Allocator, logger: Logger, input: []const u8) ![]Token {
    var tokens = ArrayList(Token).init(allocator);
    var pos: usize = 0;
    while (pos < input.len) {
        if (input[pos] == '(') {
            const token = Token{ .type = .LBRACE, .value = input[pos .. pos + 1] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 1;
            continue;
        }
        if (input[pos] == ')') {
            const token = Token{ .type = .RBRACE, .value = input[pos .. pos + 1] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 1;
            continue;
        }
        if (input[pos] == ' ') {
            const token = Token{ .type = .SPACE, .value = input[pos .. pos + 1] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 1;
            continue;
        }
        if (input[pos] == '\n') {
            const token = Token{ .type = .NEWLINE, .value = input[pos .. pos + 1] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 1;
            continue;
        }
        if (input[pos] == ',') {
            const token = Token{ .type = .COMMA, .value = input[pos .. pos + 1] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 1;
            continue;
        }
        if (pos + 3 < input.len and std.mem.eql(u8, input[pos .. pos + 3], "mul")) {
            const token = Token{ .type = .MUL, .value = input[pos .. pos + 3] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 3;
            continue;
        }
        if (pos + 4 < input.len and std.mem.eql(u8, input[pos .. pos + 4], "do()")) {
            const token = Token{ .type = .DO, .value = input[pos .. pos + 4] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 4;
            continue;
        }
        if (pos + 7 < input.len and std.mem.eql(u8, input[pos .. pos + 7], "don't()")) {
            const token = Token{ .type = .DONT, .value = input[pos .. pos + 7] };

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);

            try tokens.append(token);
            pos += 7;
            continue;
        }
        if (input[pos] <= '9' and input[pos] >= '0') {
            const res = parseNumber(input, pos);
            const token = res[0];

            const tokString = try token.String(allocator);
            logger.Debug("Detected token {s} on pos: {}", .{ tokString, pos });
            allocator.free(tokString);
            pos = res[1];

            try tokens.append(token);
            continue;
        }
        logger.DebugErr("Detected unexpected character {c} on pos: {}", .{ input[pos], pos });
        const token = Token{ .type = .ILLEGAL, .value = input[pos .. pos + 1] };
        try tokens.append(token);

        pos += 1;
    }

    return try tokens.toOwnedSlice();
}

fn solve(allocator: std.mem.Allocator, logger: Logger, tokens: []Token) !i64 {
    logger.Debug("\nStarted parsing\n", .{});
    const validSequence = [_]TokenTyp{ .MUL, .LBRACE, .NUMBER, .COMMA, .NUMBER, .RBRACE };

    var total: i64 = 0;
    var pos: usize = 0;
    var opEnabled: bool = true;
    while (pos <= tokens.len - 1) {
        if (tokens[pos].type == .MUL) {
            logger.Debug("Find possible valid sequence, pos: {}", .{pos});
            var firstNumber: ?i32 = null;
            var secondNumber: ?i32 = null;
            for (validSequence) |tok| {
                const tokString = try tokens[pos].String(allocator);
                if (tok != tokens[pos].type) {
                    logger.DebugErr("Should be {any}, got {any}", .{ tok, tokens[pos].type });
                    firstNumber = null;
                    secondNumber = null;
                    break;
                }
                logger.Debug("Token {s} on pos: {}", .{ tokString, pos });
                allocator.free(tokString);

                if (tok == .NUMBER) {
                    if (firstNumber == null) {
                        firstNumber = try std.fmt.parseInt(i32, tokens[pos].value, 10);
                    } else if (secondNumber == null) {
                        secondNumber = try std.fmt.parseInt(i32, tokens[pos].value, 10);
                    }
                }
                pos += 1;
            }
            if (firstNumber != null and secondNumber != null and opEnabled) {
                logger.Debug("mul({},{})", .{ firstNumber.?, secondNumber.? });
                total += firstNumber.? * secondNumber.?;
            }
        } else {
            if (tokens[pos].type == .DO) {
                opEnabled = true;
            } else if (tokens[pos].type == .DONT) {
                opEnabled = false;
            }
            pos += 1;
        }
    }

    return total;
}

const Level = @import("std").log.Level;
pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const Log = try NewLogger();

    Log.Debug("Input: {s}\n", .{inputStr});
    const tokens = try tokenize(allocator, Log, inputStr);
    const res = try solve(allocator, Log, tokens);

    Log.Info("{}", .{res});
}
