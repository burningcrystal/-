function ret = Code(lenchrom,bound)
flag = 0;  %用来监测生成出来的个体是否合适。
while flag==0
    pick = rand(1,length(lenchrom));   %生成0-1间的随机数
    %随机数生成:区间长度*随机的百分比（上面的pick）+区间基址
    ret = bound(:,1)'+(bound(:,2)-bound(:,1))'.*pick;
    flag = test(lenchrom,bound,ret);  %交叉后进行检验
end

end