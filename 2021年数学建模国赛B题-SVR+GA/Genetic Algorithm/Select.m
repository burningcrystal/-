function ret = Select(individuals,sizepop)
    individuals.fitness = 1./individuals.fitness;  %吧目标函数的值转化为适应函数的值

    sumfitness = sum(individuals.fitness);
    sumf = individuals.fitness./ sumfitness;   %每个个体的数占总值的百分比

    index = [];
    for i=1:sizepop
       pick = rand;
       while pick==0
           pick = rand;
       end
        for j = 1:sizepop
            pick = pick-sumf(j);
            if pick < 0
                index = [index,j];
                break;
            end
        end
    end
    
    individuals.chrom = individuals.chrom(index, :);
    individuals.fitness = individuals.fitness(index);
    ret = individuals;
end