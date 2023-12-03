#include <fstream>
#include <iostream>
#include <string>
#include <unordered_set>

using namespace std;

vector<string> FileOpener() {
  ifstream input_stream("data.txt");

  // check stream status
  if (!input_stream)
    cerr << "Can't open input file!";

  // file contents
  vector<string> text;

  // one line
  string line;

  // extract all the text from the input file
  while (getline(input_stream, line)) {

    // store each line in the vector
    text.push_back(line);
  }
  return text;
}

int getGears(vector<string> &lines, int i, int j, int mod1, int mod2,
             int &num_count) {
  string num1;
  int k = i;
  int l = j;
  bool start = false;
  if (((k + mod1) < 0) || (k + mod1 > lines.size())) {
    mod1 = 0;
  }
  if (((l + mod2) < 0) || (l + mod2 > lines[0].size())) {
    mod2 = 0;
  }
  while (lines[k + mod1][l + mod2] >= '0' && lines[k + mod1][l + mod2] <= '9') {
    l--;
    start = true;
  }
  if (start)
    l++;
  while (lines[k + mod1][l + mod2] >= '0' && lines[k + mod1][l + mod2] <= '9') {
    num1 += lines[k + mod1][l + mod2];
    lines[k + mod1][l + mod2] = '.';
    l++;
  }

  if (num1.size() == 0) {
    return 1;
  }
  cout << num1 << " scanned" << endl;
  num_count++;
  return stoi(num1);
}

int getStuff(vector<string> &lines, int i, int j, int mod1, int mod2) {
  string num1;
  int k = i;
  int l = j;
  bool start = false;
  if (((k + mod1) < 0) || (k + mod1 > lines.size())) {
    mod1 = 0;
  }
  if (((l + mod2) < 0) || (l + mod2 > lines[0].size())) {
    mod2 = 0;
  }
  while (lines[k + mod1][l + mod2] >= '0' && lines[k + mod1][l + mod2] <= '9') {
    l--;
    start = true;
  }
  if (start)
    l++;
  while (lines[k + mod1][l + mod2] >= '0' && lines[k + mod1][l + mod2] <= '9') {
    num1 += lines[k + mod1][l + mod2];
    lines[k + mod1][l + mod2] = '.';
    l++;
  }

  if (num1.size() == 0) {
    return 0;
  }
  cout << num1 << " scanned" << endl;
  return stoi(num1);
}

void solve1(vector<string> lines) {
  unordered_set<char> m{'.', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'};
  bool engine;
  int res = 0;
  for (size_t i = 0; i < lines.size(); i++) {
    for (size_t j = 0; j < lines.size(); j++) {
      engine = !m.count(lines[i][j]);
      bool is_num = false;
      // -1 -1; -1 0; 1 1;
      // 0 -1; 0 0; 0 1;
      // 1 -1; 1 0; 1 1;
      int tmp = 0;
      if (engine) {
        tmp += getStuff(lines, i, j, -1, -1);
        tmp += getStuff(lines, i, j, -1, 0);
        tmp += getStuff(lines, i, j, -1, 1);
        tmp += getStuff(lines, i, j, 0, -1);
        tmp += getStuff(lines, i, j, 0, 1);
        tmp += getStuff(lines, i, j, 1, -1);
        tmp += getStuff(lines, i, j, 1, 0);
        tmp += getStuff(lines, i, j, 1, 1);
        cout << lines[i][j] << endl;
      }
      res += tmp;
    }
  }
  cout << res << endl;
}

void solve2(vector<string> lines) {
  unordered_set<char> m{'8', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'};
  bool gear;
  int res = 0;
  for (size_t i = 0; i < lines.size(); i++) {
    for (size_t j = 0; j < lines.size(); j++) {
      int num_count = 0;
      gear = lines[i][j] == '*';
      // -1 -1; -1 0; 1 1;
      // 0 -1; 0 0; 0 1;
      // 1 -1; 1 0; 1 1;
      int tmp = 1;
      if (gear) {
        tmp *= getGears(lines, i, j, -1, -1, num_count);
        tmp *= getGears(lines, i, j, -1, 0, num_count);
        tmp *= getGears(lines, i, j, -1, 1, num_count);
        tmp *= getGears(lines, i, j, 0, -1, num_count);
        tmp *= getGears(lines, i, j, 0, 1, num_count);
        tmp *= getGears(lines, i, j, 1, -1, num_count);
        tmp *= getGears(lines, i, j, 1, 0, num_count);
        tmp *= getGears(lines, i, j, 1, 1, num_count);
        cout << lines[i][j] << endl;
      }
      if (num_count <= 1) {
        continue;
      }
      res += tmp;
    }
  }
  cout << res << endl;
}

int main() {
  vector<string> lines;
  lines = FileOpener();
  vector<string> test_lines;
  test_lines.push_back("467..114..");
  test_lines.push_back("...*......");
  test_lines.push_back("..35..633.");
  test_lines.push_back("......#...");
  test_lines.push_back("617*......");
  test_lines.push_back(".....+.58.");
  test_lines.push_back("..592.....");
  test_lines.push_back("......755.");
  test_lines.push_back("...$.*....");
  test_lines.push_back(".664.598..");
  solve1(lines);
  solve2(lines);
}
