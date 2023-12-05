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
