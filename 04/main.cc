#include "../cc_common/common.h"
#include <algorithm>
#include <iostream>
#include <set>
#include <stack>
#include <stdexcept>
#include <string>
#include <unordered_map>
#include <unordered_set>

using namespace std;

bool assertWin(int match_param, int num, unordered_set<int> winning_numbers) {
  if (!match_param) {
    return false;
  }
  return winning_numbers.count(num) > 0;
}

int getCardResult(std::string result) {
  result = result.substr(result.find(":"));
  result.erase(result.begin());
  result.erase(result.begin());

  unordered_set<int> winning_numbers;

  size_t idx = 0;
  int match_param = 0;
  int win_points = 0;

  for (size_t i = 0; i < result.size(); i++) {
    string substr = result.substr(i);
    try {
      int n = stoi(substr, &idx, 10);
      i += idx;
      if (!match_param) {
        winning_numbers.insert(n);
      }
      if (assertWin(match_param, n, winning_numbers)) {
        if (win_points == 0) {
          win_points = 1;
        } else {
          win_points *= 2;
        }
      }
    } catch (std::invalid_argument) {
      match_param = 1;
    }
  }
  return win_points;
}

void solve1(vector<string> lines) {
  int points = 0;
  for (size_t i = 0; i < lines.size(); i++) {
    points += getCardResult(lines[i]);
  }
  cout << points << endl;
}

struct card {
  int number;
  unordered_set<int> winning_numbers;
  vector<int> scratched_numbers;
  card(){

  };
  card(int n) : number(n){};
  card(int n, unordered_set<int> w, vector<int> s)
      : number(n), winning_numbers(w), scratched_numbers(s){};

  card(const card &other)
      : number(other.number), winning_numbers(other.winning_numbers),
        scratched_numbers(other.scratched_numbers) {}

  // Move constructor
  card(card &&other) noexcept
      : number(std::move(other.number)),
        winning_numbers(std::move(other.winning_numbers)),
        scratched_numbers(std::move(other.scratched_numbers)) {}

  // Copy assignment operator
  card &operator=(const card &other) {
    if (this != &other) {
      number = other.number;
      winning_numbers = other.winning_numbers;
      scratched_numbers = other.scratched_numbers;
    }
    return *this;
  }

  // Move assignment operator
  card &operator=(card &&other) noexcept {
    if (this != &other) {
      number = std::move(other.number);
      winning_numbers = std::move(other.winning_numbers);
      scratched_numbers = std::move(other.scratched_numbers);
    }
    return *this;
  }

  bool operator==(const card &r) const { return this->number == r.number; }
  bool operator<(const card &r) const { return this->number > r.number; }
  bool operator>(const card &r) const { return this->number < r.number; }
};

struct cardHash {
  std::size_t operator()(const card &c) const {
    return std::hash<int>()(c.number);
  }
};

card getCardCopy(std::string &result) {
  result = result.substr(result.find("d"));
  result.erase(result.begin());
  size_t idx = 0;
  int cardNum = stoi(result, &idx);
  result = result.substr(idx + 2);
  int match_param = 0;
  unordered_set<int> winning_numbers;
  vector<int> scratched_numbers;

  for (size_t i = 0; i < result.size(); i++) {
    string substr = result.substr(i);
    try {
      int n = stoi(substr, &idx, 10);
      i += idx;
      if (!match_param) {
        winning_numbers.insert(n);
      } else {
        scratched_numbers.push_back(n);
      }
    } catch (std::invalid_argument) {
      match_param = 1;
    }
  }

  card n = card(cardNum, winning_numbers, scratched_numbers);

  return n;
}

vector<int> getCardCopyCount(const card &c) {
  int count = 0;
  for (const int &n : c.scratched_numbers) {
    count += c.winning_numbers.count(n) > 0 ? 1 : 0;
  }
  vector<int> copies;
  for (int i = 0; i < count; i++) {
    copies.push_back(c.number + i + 1);
  }
  return copies;
}

void solve2(vector<string> &lines) {
  vector<card> cards;
  for (size_t i = 0; i < lines.size(); i++) {
    cards.push_back(getCardCopy(lines[i]));
  }

  stack<card> s;
  for (auto it = cards.rbegin(); it != cards.rend(); it++) {
    s.push(*it);
  }
  int count = 0;
  while (!s.empty()) {
    card c = s.top();
    s.pop();
    vector<int> cards_nums = getCardCopyCount(c);
    for (size_t i = 0; i < cards_nums.size(); i++) {
      s.push(cards[cards_nums[i] - 1]);
    }
    count++;
  }
  cout << count << endl;
}

int maxCard = 0;

void solve2_1(vector<string> &lines) {
  vector<card> cards(lines.size());
  unordered_map<card, int, cardHash> v;
  for (size_t i = 0; i < lines.size(); i++) {
    card c = getCardCopy(lines[i]);
    cards[i] = c;
    maxCard = max(maxCard, c.number);
    v[c] = 1;
  }
  for (int i = 1; i <= maxCard; i++) {
    int value = v[i];
    auto key = cards[i - 1];
    for (int j = 0; j < value; j++) {
      auto vv = getCardCopyCount(key);
      for (size_t i = 0; i < vv.size(); i++) {
        v[vv[i]]++;
      }
    }
  }
  int count = 0;
  for (const auto &[key, value] : v) {
    count += value;
  }
  cout << count << endl;
}

int main() {
  vector<string> lines = FileOpener();
  vector<string> test_lines;

  test_lines.push_back("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53");
  test_lines.push_back("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19");
  test_lines.push_back("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1");
  test_lines.push_back("Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83");
  test_lines.push_back("Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36");
  test_lines.push_back("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11");
  solve1(lines);
  solve2_1(lines);
  // solve2(lines);
}
