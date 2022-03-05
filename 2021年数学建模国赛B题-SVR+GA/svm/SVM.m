 M1 = fitrsvm(smp,label,'Standardize',true);
 y_pre = predict(M1,smp_tst);

 