// Minimal ambient declarations for the Node APIs the theme's unit tests use.
//
// The theme doesn't depend on @types/node, so tsconfig.test.json sets
// "types": [] and this file supplies only what the tests actually touch.

declare const __dirname: string;

declare const require: {
    (id: string): any;
    resolve(id: string): string;
    cache: { [id: string]: any };
};

declare module "node:test" {
    export function test(name: string, fn: () => void | Promise<void>): void;
}

declare module "assert" {
    function ok(value: any, message?: string): void;
    function strictEqual(actual: any, expected: any, message?: string): void;
    function notStrictEqual(actual: any, expected: any, message?: string): void;
    function deepStrictEqual(actual: any, expected: any, message?: string): void;
    function match(value: string, regExp: RegExp, message?: string): void;
    function doesNotMatch(value: string, regExp: RegExp, message?: string): void;
}
