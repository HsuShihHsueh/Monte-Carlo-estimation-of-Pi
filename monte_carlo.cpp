#include <iostream>
#include <cstdlib>
#include <ctime> 
#include <cmath>
#define n 1'000'000 //separators started in C++14
//#define n 1000000 //use it before C++14

double rand_unit_interval(){
    double tmp = (double) rand() / (RAND_MAX + 1.0);
    return tmp;
}

double average_value_method(){
    double pi = 0;
    for(int i=0;i<n;i++){
        double tmp = rand_unit_interval();
        tmp = 1 - pow(tmp,2);
        tmp = sqrt(tmp);
        pi += tmp;
    }
    pi = 4 * pi / n;
    return pi;
}

double area_method(){
    double pi = 0;
    for(int i=0;i<n;i++){
        double tmp = pow(rand_unit_interval(),2) + pow(rand_unit_interval(),2);
        if(tmp<1){
            pi++;
        }
    }    
    pi = 4 * pi / n;
    return pi;
}

int main(){
    srand( time(NULL) );
    double result_average = average_value_method();
    std::cout << "The estimate result of average value method: " << result_average << std::endl;
    double result_area = area_method();
    std::cout << "The estimate result of area method: " << result_area << std::endl;    
    return 0;
}
