maxgen = 300;  %迭代次数
PopSize = 100;  %种群大小
Pcross = 0.6;    %交叉率
pmutation = 0.01;  %变异率
k=1
figure




%目标函数  
%变量：
lenchrom = [1,1,1,1,1];  %变量个数的实数编码
bound = [10,200; 0.5,5 ; 10,200 ; 0.3,2.1 ; 250,400];  %每种变量的范围
% 我们的目的是求函数的最值和相关参数的值
geti = struct('fitness',zeros(1,PopSize) , 'chrom',[]); %定义适应度（一百个个体，一百个适应度）


bestfitness = [];  %目标
bestchrom = [];   %目标

%初始化种群：
for i = 1:PopSize
    geti.chrom(i,:) = Code(lenchrom,bound);     %code函数是生成范围内随机数的函数
    x=  geti.chrom(i,:);
    geti.fitness(i) = fun(x,M1);
end

[bestfitness,bestindex] = max(geti.fitness);  %找到当前的最优解
bestchrom = geti.chrom(bestindex,:);
trace = [trace;bestfitness];

subplot(3,3,k);
scatter(1:100,geti.fitness);
k=k+1;

%开始进入大循环：选择、交叉、变异、择优

for i=(1:maxgen)
    geti = Select(geti,PopSize);  %进行选择种群的函数
    avgfitness = sum(geti.fitness) / PopSize;
    geti.chrom = Cross(Pcross,lenchrom,geti.chrom,PopSize,bound); %交叉
    geti.chrom = Mutation(pmutation,lenchrom,geti.chrom,PopSize,[i,maxgen],bound); %变异
    
    for j= 1:PopSize    %更新目标值
        x=  geti.chrom(j,:);
        geti.fitness(j) = fun(x,M1);
    end
    
    [newbestfitness,newbestindex] =max(geti.fitness);
    [worsefitness , worseindex] =min(geti.fitness);
    if bestfitness < newbestfitness         %更新最优目标值
        bestfitness = newbestfitness;
        bestchrom = geti.chrom(bestindex,:);
    end
    
    geti.chrom(worseindex,:) = bestchrom;
    geti.fitness(worseindex) = bestfitness;
    
    if i==5||i==50||i==150||i==250
        subplot(3,3,k)
        scatter(1:100,geti.fitness);
        k=k+1
    end

        
    
    trace = [trace;bestfitness];
end


