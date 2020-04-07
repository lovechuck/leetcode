package src

/*
嵌套列表权重和
题目
给定一个嵌套的整数列表，请返回该列表按深度加权后所有整数的总和。

每个元素要么是整数，要么是列表。同时，列表中元素同样也可以是整数或者是另一个列表。

示例 1:

输入: [[1,1],2,[1,1]]
输出: 10
解释: 因为列表中有四个深度为 2 的 1 ，和一个深度为 1 的 2。
示例 2:

输入: [1,[4,[6]]]
输出: 27
解释: 一个深度为 1 的 1，一个深度为 2 的 4，一个深度为 3 的 6。所以，1 + 4*2 + 6*3 =
27
*/

/**
 * public interface NestedInteger {
 *     public NestedInteger();//无参构造方法
 *     public NestedInteger(int value);//构造方法
 *     public boolean isInteger();//判断此对象是否为Integer类型
 *     public Integer getInteger();//返回其中的Integer类型值
 *     public void setInteger(int value);//修改对象中的值
 *     public void add(NestedInteger ni);//为此对象中添加一个新的对象
 *     public List<NestedInteger> getList();//返回其中的list类型
 * }
 */

/*
public int depthSum(List<NestedInteger> nestedList)
{
	int x =help(nestedList,1);
	return x;
}
public int help(List<NestedInteger>n,int lev)
{
	int sum=0;
	for(NestedInteger x: n)
	{
		if(x.isInteger())
		{
			sum+=x.getInteger()*lev;
		}
		else
		{
			sum+=help(x.getList(),lev+1);
		}
	}
	return sum;
}
*/
