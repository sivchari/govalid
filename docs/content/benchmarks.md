---
title: "Benchmarks"
description: "Performance comparison between govalid and go-playground/validator"
weight: 30
---

# Performance Benchmarks

govalid is designed for maximum performance with zero allocations. Here are the latest benchmark results comparing govalid with go-playground/validator.

## Latest Results

**Benchmarked on:** 2025-07-16  
**Platform:** Darwin 24.5.0 arm64  
**Go version:** go1.25-devel_a8e99ab19c

## Raw Benchmark Data

```
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08221 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08211 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08386 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08506 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08340 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08368 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08714 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08768 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08579 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELConcurrent-16             	1000000000	         0.08608 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8317 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8408 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8364 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8303 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8365 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8298 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8397 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8467 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8313 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-16    	1000000000	         0.8305 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09895 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.1004 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09969 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09850 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.1153 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09841 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09825 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09865 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09827 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCacheEffectiveness-16     	1000000000	         0.09830 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	619130752	         1.944 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	625241518	         1.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	630330586	         1.911 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	625552111	         1.919 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	627390291	         1.956 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	617119010	         1.951 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	624382257	         1.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	614515627	         1.928 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	626190958	         1.938 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-16                  	623657806	         1.943 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	618552051	         1.923 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	628656386	         1.908 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	634784330	         1.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	620215537	         1.960 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	613909132	         1.937 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	620773004	         1.930 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	627497324	         1.926 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	621488347	         1.943 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	607611730	         1.951 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-16             	629428676	         1.912 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	649483614	         1.899 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	614503958	         1.948 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	619591082	         1.905 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	648771241	         1.872 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	627424734	         1.886 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	612057661	         1.950 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	632546076	         1.950 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	590049792	         1.928 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	627348062	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-16           	644056698	         1.923 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	625452256	         1.952 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	640366999	         1.889 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	640784460	         1.899 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	631436182	         1.936 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	629614714	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	621101535	         1.921 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	630729951	         1.896 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	641576278	         1.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	613408994	         1.939 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-16      	608121462	         1.923 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	32320837	        37.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	32522931	        36.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31914044	        37.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31555072	        37.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	32476831	        37.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	33073188	        36.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31532167	        37.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	32599246	        37.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31843753	        36.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-16                     	31661364	        37.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1887846	       634.5 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1883122	       633.4 ns/op	      90 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1906246	       626.0 ns/op	      90 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1928503	       621.0 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1910904	       628.6 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1868458	       634.7 ns/op	      90 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1929516	       621.0 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1931611	       623.0 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1934497	       625.7 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoPlaygroundEmail-16                	 1904106	       633.5 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-16                 	 2045612	       586.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2044362	       586.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2051377	       588.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2017550	       595.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2052342	       585.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2050610	       586.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2044392	       591.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2027974	       594.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2027649	       588.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorEmail-16                 	 2047588	       585.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationEmail-16              	28794844	        42.08 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	27640428	        42.77 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	27754327	        42.66 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28523097	        41.92 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28902517	        41.63 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28885297	        42.90 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28203301	        42.71 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28913427	        42.22 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28178826	        41.98 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationEmail-16              	28817088	        42.38 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateEmail-16              	  110956	     10703 ns/op	   16101 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  110484	     10697 ns/op	   16124 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  112976	     10637 ns/op	   16102 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  113974	     10415 ns/op	   16134 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  112286	     10595 ns/op	   16106 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  112124	     10737 ns/op	   16101 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  111264	     10700 ns/op	   16128 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  113616	     10541 ns/op	   16125 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  117296	     10617 ns/op	   16112 B/op	      74 allocs/op
BenchmarkGookitValidateEmail-16              	  111259	     10674 ns/op	   16141 B/op	      74 allocs/op
BenchmarkGoValidEnum-16                      	545223546	         2.203 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	538649304	         2.219 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	538552504	         2.234 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	537160402	         2.199 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	549945033	         2.213 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	544636765	         2.212 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	540510105	         2.232 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	539215980	         2.212 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	542151934	         2.208 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEnum-16                      	530166570	         2.227 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	609233181	         1.939 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	609520455	         1.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	642166810	         1.880 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	634752572	         1.872 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	639555864	         1.926 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	607742895	         1.943 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	639840182	         1.896 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	637850866	         1.881 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	637952173	         1.927 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-16                        	639243561	         1.933 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19218500	        61.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19425006	        61.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19064388	        62.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19268002	        62.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19769940	        61.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19435440	        61.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19506273	        62.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19097700	        63.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19001346	        61.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-16                   	19116003	        62.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21508635	        54.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21594757	        54.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22058874	        54.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22456314	        54.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21988878	        54.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	21308992	        54.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22320164	        55.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22036273	        54.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22219443	        54.39 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-16                    	22310828	        54.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	617906678	         1.932 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	629888472	         1.890 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	637901866	         1.864 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	638075988	         1.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	616244592	         1.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	624969076	         1.878 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	644921611	         1.875 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	650839218	         1.901 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	627020127	         1.938 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-16                       	638061283	         1.888 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19634577	        60.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19754604	        60.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19926823	        61.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19524086	        61.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19482416	        61.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19885574	        60.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19517523	        61.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19558812	        61.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19732311	        61.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-16                  	19253292	        61.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	614953226	         1.941 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	613055265	         1.940 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	630986290	         1.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	633449844	         1.903 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	627177835	         1.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	615834514	         1.933 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	632111110	         1.895 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	634569213	         1.906 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	640230765	         1.931 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLT-16                        	616444029	         1.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19889241	        60.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19531964	        60.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19983194	        60.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19521810	        61.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19769900	        60.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19840216	        60.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19751716	        60.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19307192	        61.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19343709	        60.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-16                   	19628583	        60.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	626405058	         1.911 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	617536888	         1.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	611867684	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	632066026	         1.892 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	636954792	         1.907 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	609850094	         1.950 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	612017601	         1.925 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	634074922	         1.898 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	641211602	         1.874 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-16                       	644934030	         1.927 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	18794943	        61.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19611901	        60.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19959609	        60.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19519190	        61.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19832470	        61.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19741022	        60.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19819668	        60.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19816599	        61.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19229073	        61.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-16                  	19970100	        60.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	483567601	         2.471 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	486643740	         2.487 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	478956751	         2.506 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	477148162	         2.490 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	488703649	         2.462 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	468681270	         2.500 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	477940315	         2.508 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	477892255	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	486582158	         2.477 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-16                  	462870027	         2.514 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	14777430	        79.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15075565	        79.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15433611	        78.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15363908	        78.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15203986	        79.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	14829880	        80.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15025762	        78.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15236515	        78.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15416749	        79.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-16             	15387171	        79.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	77349903	        15.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	77077270	        15.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	78205618	        15.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	77117516	        15.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	75641772	        15.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	78862191	        15.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	79146754	        15.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	77339520	        15.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	75630054	        15.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-16                 	79458358	        15.37 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16548746	        72.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16525946	        72.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16437445	        72.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16783422	        71.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16861362	        71.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16246596	        72.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16487479	        72.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16709174	        71.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16845394	        72.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-16            	16810371	        72.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7355149	       161.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7456647	       158.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7524546	       159.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7547952	       162.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7294558	       163.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7522123	       158.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7522021	       158.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7585036	       160.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7371297	       162.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMaxLength-16             	 7321084	       160.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6876850	       176.7 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6719857	       179.3 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6707959	       181.4 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6632403	       179.9 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6665997	       178.7 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6856638	       177.4 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6786402	       176.7 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6874554	       177.1 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6928856	       173.7 ns/op	     432 B/op	       4 allocs/op
BenchmarkOzzoValidationMaxLength-16          	 6387793	       183.0 ns/op	     432 B/op	       4 allocs/op
BenchmarkGookitValidateMaxLength-16          	  114445	     10133 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  115782	     10214 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  117318	     10289 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  116253	     10373 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  115336	     10294 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  115688	     10310 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  117314	     10131 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  118848	     10227 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  118149	     10242 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGookitValidateMaxLength-16          	  117813	     10207 ns/op	   15632 B/op	      80 allocs/op
BenchmarkGoValidMinItems-16                  	432711421	         2.755 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	436474869	         2.756 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	437049604	         2.764 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	433022473	         2.790 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	433878693	         2.754 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	436022151	         2.750 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	435682058	         2.759 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	431841016	         2.795 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	430346542	         2.765 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinItems-16                  	438246168	         2.747 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15480092	        78.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15002007	        79.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15097773	        79.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15097338	        78.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15146685	        78.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15097662	        79.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	14933901	        79.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15048704	        79.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15250448	        79.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-16             	15352074	        79.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-16                 	100000000	        11.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17845170	        66.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17795066	        67.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17794384	        67.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17491627	        67.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17856655	        66.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17850336	        67.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	18060886	        67.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17569149	        67.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17935700	        66.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-16            	17992914	        67.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-16             	 7542423	       161.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7509517	       163.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7561120	       158.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7603398	       159.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7479520	       163.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7481840	       161.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7282023	       161.8 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7557021	       159.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7526383	       161.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidatorMinLength-16             	 7408116	       163.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidRequired-16                  	613623986	         1.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	621332950	         1.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	628562534	         1.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	613784314	         1.946 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	617220583	         1.932 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	634786849	         1.901 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	633633534	         1.910 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	611511448	         1.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	616228636	         1.930 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidRequired-16                  	632464813	         1.904 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14503328	        83.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14340872	        85.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14117618	        84.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14378016	        83.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14159347	        83.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14145681	        84.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14431640	        84.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14451178	        83.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14502846	        83.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-16             	14393688	        84.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	626280423	         1.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	636268496	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	637125571	         1.887 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	654244950	         1.892 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	621867720	         1.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	601055979	         1.896 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	649260912	         1.871 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	647959534	         1.896 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	617009406	         1.943 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-16              	603587320	         1.877 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationRequired-16           	36211656	        33.34 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	36598988	        33.89 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	34378740	        34.70 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	34908201	        34.42 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	36202098	        33.43 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	35598140	        33.99 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	35508912	        34.46 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	34140171	        34.60 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	36065503	        33.69 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationRequired-16           	35759383	        34.29 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateRequired-16           	  118464	     10012 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  117220	     10201 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  117026	     10288 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  118225	     10218 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  117763	     10136 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  121041	     10144 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  118038	     10318 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  115998	     10265 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  117462	     10060 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGookitValidateRequired-16           	  122958	     10272 ns/op	   15472 B/op	      72 allocs/op
BenchmarkGoValidURL-16                       	28069600	        42.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28339009	        42.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28746961	        42.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28230339	        42.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28121594	        42.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	29197582	        41.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28424426	        42.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28450569	        42.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	27898866	        42.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidURL-16                       	28550016	        41.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4308352	       280.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4122770	       287.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4122933	       285.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4231399	       281.3 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4311073	       282.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4201960	       287.0 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4120405	       290.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4298881	       279.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4297384	       282.1 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoPlaygroundURL-16                  	 4213228	       286.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  155608	      7807 ns/op	     146 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  158746	      7676 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  157711	      7687 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  155072	      7796 ns/op	     146 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  156248	      7819 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  154514	      7713 ns/op	     146 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  153856	      7730 ns/op	     146 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  157533	      7728 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  155770	      7795 ns/op	     148 B/op	       1 allocs/op
BenchmarkGoValidatorURL-16                   	  154666	      7735 ns/op	     145 B/op	       1 allocs/op
BenchmarkOzzoValidationURL-16                	  154232	      7718 ns/op	     171 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  156506	      7783 ns/op	     169 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  155185	      7864 ns/op	     171 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  151880	      7844 ns/op	     169 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  151316	      7789 ns/op	     169 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  154465	      7911 ns/op	     171 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  154320	      7835 ns/op	     170 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  154112	      7768 ns/op	     172 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  156765	      7730 ns/op	     169 B/op	       2 allocs/op
BenchmarkOzzoValidationURL-16                	  151394	      7866 ns/op	     169 B/op	       2 allocs/op
BenchmarkGookitValidateURL-16                	  114381	     10589 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  113596	     10750 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  109780	     10736 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  112147	     10786 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  110826	     10670 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  110617	     10908 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  111747	     10805 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  110395	     10785 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  110246	     10552 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGookitValidateURL-16                	  112106	     10733 ns/op	   15641 B/op	      75 allocs/op
BenchmarkGoValidUUID-16                      	33149475	        36.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	33537663	        36.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	32978546	        36.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	32720504	        37.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	31934498	        36.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	32104016	        36.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	33037524	        36.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	32363331	        37.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	32000355	        36.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidUUID-16                      	33638532	        36.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4776891	       248.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4709527	       252.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4738076	       251.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4810500	       248.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4818919	       250.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4710769	       253.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4740919	       250.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4829394	       247.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4759467	       252.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-16                 	 4720359	       256.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6036775	       198.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6021385	       198.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6020737	       196.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6005136	       199.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6103957	       198.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6087932	       194.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6170476	       195.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6161251	       196.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6087652	       197.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-16                  	 6094603	       195.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkOzzoValidationUUID-16               	 5089902	       235.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 5068016	       237.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 4978867	       240.4 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 5019273	       236.0 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 5117695	       234.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 5125396	       238.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 4993830	       242.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 4925216	       244.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 4990972	       238.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkOzzoValidationUUID-16               	 4785774	       245.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkGookitValidateUUID-16               	  121904	     10371 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  113184	     10823 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  111711	     10877 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  111168	     10834 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  110126	     10645 ns/op	   15518 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  112779	     10770 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  112326	     10823 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  110032	     10826 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  110846	     10708 ns/op	   15519 B/op	      74 allocs/op
BenchmarkGookitValidateUUID-16               	  109089	     10735 ns/op	   15518 B/op	      74 allocs/op
```

## Performance Comparison

| Validator | govalid (ns/op) | go-playground/validator (ns/op) | Improvement | govalid Allocs | Competitor Allocs |
|-----------|-----------------|--------------------------------|-------------|----------------|-------------------|
| Email | 37.65ns | 634.5 | **16.9x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| Email | 37.65ns | 633.4 | **16.8x faster** | 0 allocs/op | 5 allocs + 90 B/op |
| Email | 37.65ns | 626.0 | **16.6x faster** | 0 allocs/op | 5 allocs + 90 B/op |
| Email | 37.65ns | 621.0 | **16.5x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| Email | 37.65ns | 628.6 | **16.7x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| Email | 37.65ns | 634.7 | **16.9x faster** | 0 allocs/op | 5 allocs + 90 B/op |
| Email | 37.65ns | 621.0 | **16.5x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| Email | 37.65ns | 623.0 | **16.5x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| Email | 37.65ns | 625.7 | **16.6x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| Email | 37.65ns | 633.5 | **16.8x faster** | 0 allocs/op | 5 allocs + 89 B/op |
| GT | 1.933ns | 61.90 | **32.0x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 61.95 | **32.0x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 62.49 | **32.3x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 62.87 | **32.5x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 61.35 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 61.91 | **32.0x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 62.13 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 63.14 | **32.7x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 61.97 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| GT | 1.933ns | 62.19 | **32.2x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 60.81 | **32.2x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 60.93 | **32.3x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.16 | **32.4x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.53 | **32.6x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.03 | **32.3x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 60.87 | **32.2x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.12 | **32.4x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.13 | **32.4x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.19 | **32.4x faster** | 0 allocs/op | 0 allocs/op |
| GTE | 1.888ns | 61.09 | **32.4x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.24 | **30.9x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.54 | **31.1x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.93 | **31.3x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 61.35 | **31.5x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.79 | **31.2x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.59 | **31.1x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.86 | **31.2x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 61.72 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.35 | **31.0x faster** | 0 allocs/op | 0 allocs/op |
| LT | 1.949ns | 60.24 | **30.9x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 61.16 | **31.7x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 60.72 | **31.5x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 60.18 | **31.2x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 61.47 | **31.9x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 61.66 | **32.0x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 60.68 | **31.5x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 60.95 | **31.6x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 61.33 | **31.8x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 61.91 | **32.1x faster** | 0 allocs/op | 0 allocs/op |
| LTE | 1.927ns | 60.80 | **31.6x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 79.98 | **31.8x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 79.19 | **31.5x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 78.48 | **31.2x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 78.49 | **31.2x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 79.37 | **31.6x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 80.10 | **31.9x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 78.77 | **31.3x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 78.91 | **31.4x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 79.25 | **31.5x faster** | 0 allocs/op | 0 allocs/op |
| MaxItems | 2.514ns | 79.42 | **31.6x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.16 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.57 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.44 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 71.74 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 71.45 | **4.6x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.92 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.51 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 71.53 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.02 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MaxLength | 15.37ns | 72.60 | **4.7x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 78.57 | **28.6x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.73 | **29.0x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.28 | **28.9x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 78.76 | **28.7x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 78.67 | **28.6x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.82 | **29.1x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.76 | **29.0x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.51 | **28.9x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.32 | **28.9x faster** | 0 allocs/op | 0 allocs/op |
| MinItems | 2.747ns | 79.66 | **29.0x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 66.77 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.11 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.66 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.85 | **5.9x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 66.78 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.47 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.35 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.89 | **5.9x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 66.67 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| MinLength | 11.58ns | 67.61 | **5.8x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 83.23 | **43.7x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 85.52 | **44.9x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 84.25 | **44.2x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 83.22 | **43.7x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 83.56 | **43.9x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 84.45 | **44.4x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 84.78 | **44.5x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 83.24 | **43.7x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 83.35 | **43.8x faster** | 0 allocs/op | 0 allocs/op |
| Required | 1.904ns | 84.59 | **44.4x faster** | 0 allocs/op | 0 allocs/op |
| URL | 41.87ns | 280.4 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 287.6 | **6.9x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 285.5 | **6.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 281.3 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 282.5 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 287.0 | **6.9x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 290.1 | **6.9x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 279.8 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 282.1 | **6.7x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| URL | 41.87ns | 286.5 | **6.8x faster** | 0 allocs/op | 1 allocs + 144 B/op |
| UUID | 36.18ns | 248.6 | **6.9x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 252.1 | **7.0x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 251.7 | **7.0x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 248.4 | **6.9x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 250.0 | **6.9x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 253.1 | **7.0x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 250.0 | **6.9x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 247.4 | **6.8x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 252.3 | **7.0x faster** | 0 allocs/op | 0 allocs/op |
| UUID | 36.18ns | 256.3 | **7.1x faster** | 0 allocs/op | 0 allocs/op |
| Enum | 2.242ns | N/A (govalid exclusive) | **govalid exclusive** | 0 allocs/op | N/A |

## Performance Categories

### 🚀 Ultra-Fast (< 3ns)
- **Required**: ~1.9ns - 45x faster
- **GT/GTE/LT/LTE**: ~1.9ns - 32x faster
- **Enum**: ~2.2ns - govalid exclusive
- **MaxItems**: ~2.5ns - 32x faster
- **MinItems**: ~2.8ns - 29x faster

### ⚡ Fast (3-40ns)
- **MinLength**: ~11ns - 6x faster
- **MaxLength**: ~15ns - 5x faster
- **UUID**: ~36ns - 7x faster
- **URL**: ~41ns - 7x faster
- **Email**: ~36ns - 17x faster

## Key Performance Insights

### 1. Zero Allocations
**All govalid validators perform zero heap allocations**, while competitors often allocate 0-5 objects per validation.

### 2. Sub-Nanosecond Efficiency
Simple validators (GT, LT, Required) execute in under 2ns, making them essentially free operations.

### 3. Complex Validation Optimization
Even complex validators like email and URL are optimized with:
- Manual string parsing (no regex overhead)
- Single-pass validation algorithms
- Zero memory allocations

### 4. String Length Performance
Unicode-aware string validators are 4.8-6.0x faster despite proper UTF-8 handling.

## govalid-Exclusive Features

### Enum Validation
```go
// +govalid:enum=admin,user,guest
Role string
```
- **Performance**: ~2.2ns with 0 allocations
- **No equivalent** in go-playground/validator
- Supports string, numeric, and custom types

### Extended Collection Support
```go
// +govalid:maxitems=10
Items map[string]int  // Maps supported!

// +govalid:minitems=1
Channel chan string   // Channels supported!
```

## Optimization Techniques

### 1. Code Generation
- **Compile-time validation functions** (no runtime reflection)
- **Inlined simple operations** for maximum speed
- **Direct field access** with no interface overhead

### 2. External Helper Functions
Complex validators use optimized external functions for better performance.

### 3. Manual String Parsing
- **Character-by-character parsing** instead of `strings.Split`
- **Direct indexing** instead of `strings.Contains`
- **Single-pass algorithms** for complex validation

### 4. Memory Optimization
- **Zero heap allocations** across all validators
- **Stack-only operations** for maximum cache efficiency
- **Minimal memory footprint** in generated code

## Running Benchmarks

To run benchmarks yourself:

```bash
# Sync all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem
```

## Conclusion

govalid delivers exceptional performance improvements:
- **4.8x to 45x faster** than go-playground/validator
- **Zero allocations** across all validators
- **Sub-3ns performance** for simple operations
- **Extended type support** (maps, channels, enums)
- **Production-ready** with comprehensive test coverage

Choose govalid when performance matters and zero allocations are critical for your application's success.
