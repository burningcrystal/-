function flag = test(lenchrom,bound,code)   %监测一下交叉后是否出界
flag = 1;
[n, ~] = size(code);
for i = 1:n
    if code(i) < bound(i,1) ||code(i)>bound(i,2)
        flag=0;
    end
end
end