
## 反转整数

给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321

示例 2:

输入: -123
输出: -321

示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 \[−2的31次方,  2的31次方 − 1\]。根据这个假设，如果反转后的整数溢出，则返回 0。


```

//pop operation:
pop = x % 10;
x /= 10;

//push operation:
temp = rev * 10 + pop; //这句可能会导致溢出   pop > 7 会导致溢出
rev = temp;

因为2的31次方到 2的31次方-1

最大值是2147483647
最小值-2147483648

所以尾数大于7会溢出

```




### 提示

[https://leetcode-cn.com/articles/reverse-integer/](https://leetcode-cn.com/articles/reverse-integer/)
