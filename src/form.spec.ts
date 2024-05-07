import { testMe } from "form";

describe("Our Test cases", function () {
  it("Test 1", function () {
    expect(testMe(4, 2)).toBe(8);
  });

  it("Test 2", function () {
    expect(testMe(2, 10)).toBe(20);
  });
});
