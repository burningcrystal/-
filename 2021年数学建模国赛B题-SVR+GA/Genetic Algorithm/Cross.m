function ret = Cross(pcross,lenchrom,chrom,sizepop,bound)
for i = 1:sizepop
   pick = rand(1,2);  %要对两个问题进行交叉，所以生成两个随机数
   while prod(pick)==0
       pick=rand(1,2);
   end
   index = ceil(pick.*sizepop);  %向上取整，也就得到了要交叉的两个个体的idx
   pick = rand;
   
   while pick ==0
       pick = rand;
   end
   
   if pick>pcross
       continue;
   end
   
   flag=0;
   while flag==0
       pick = rand;
       while pick == 0;
           pick = rand;
       end
       pos = ceil(pick.*sum(lenchrom));   %选择出来要交叉的染色体位点
       pick = rand;
       v1 = chrom(index(1),pos);
       v2 = chrom(index(2),pos);
       
       chrom(index(1),pos)=pick*v2 + (1-pick)*v1;
       chrom(index(2),pos) = pick * v1 +(1-pick)*v2;
       flag1 = test(lenchrom,bound,chrom(index(1),:));
       flag2 = test(lenchrom,bound,chrom(index(2),:));
       
       if flag1*flag2 ==0
           flag=0;
       else
           flag=1;
       end
   end
end
ret = chrom;           
end