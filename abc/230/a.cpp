#include <iostream>
#include <iomanip>
using namespace std;

int main()
{
    int N;
    cin >> N;
    if (N >= 42) cout << "AGC" << setw(3) << setfill('0') << N+1 << endl;
    else cout << "AGC" << setw(3) << setfill('0') << N << endl;
    return 0;
}
