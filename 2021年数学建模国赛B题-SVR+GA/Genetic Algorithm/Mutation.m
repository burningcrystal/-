function ret = Mutation(pmutation , lenchrom , chrom , sizepop, pop , bound)
for i = 1:sizepop
    pick = rand;
    while pick == 0
        pick = rand;
    end
    index = ceil(pick*sizepop);   %找到突变的个体
    pick = rand;
    if pick  > pmutation
        continue;
    end
    flag = 0;
    while flag ==0
        pick = rand;
        while pick == 0
            pick = rand;
        end
        pos = ceil(pick*sum(lenchrom));   %找到突变的基因位点
        v = chrom(i,pos);
        v1 = v-bound(pos,1);
        v2 =  bound(pos,2)-v;
        pick = rand;
        if pick >0.5
            delta = v2*(1-pick ^ ((1-pop(1) / pop(2)) ^2 ));
            chrom(i,pos) = v+delta;
        else
            delta = v1*(1-pick ^ ((1-pop(1) / pop(2)) ^2 ));
            chrom(i,pos) = v-delta;
        end
        flag = test(lenchrom , bound , chrom(i,:));
    end
end
ret = chrom;
end