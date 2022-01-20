#include <algorithm>
#include <array>
#include <iostream>
using namespace std;

int main()
{
    int N;
    int score = 0;
    cin >> N;
    array<int, 34> A{};

    for (int i = 0; i < N; i++) {
        cin >> A[i];
    }

    sort(A.rbegin(), A.rend());
    for (int i = 0; i < N; i++) {
        if ((A[i] - 1) != A[i+1])
            score += A[i];
    }

    cout << score << endl;
    return 0;
}
