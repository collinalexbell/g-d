#include <stdio.h>
#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main(int argc, char** argv) {
	string cmd;
	if(argc >= 2){
		cout << argv[1] << endl;
		cmd = argv[1];
	}
	if(cmd == "list" || cmd == "l"){
		string line;
		ifstream f ("/home/kuberlog/tasks.txt");
		if(f.is_open()) {
			while (getline(f, line)){
				cout << line << '\n';
			}
			f.close();
		}
		else {
			cout << "unable to open tasks.txt";
		}
	}
	if(cmd == "add" || cmd == "a") {
		ofstream f ("/home/kuberlog/tasks.txt", std::ofstream::app);
		if(f.is_open()) {
			for(int i = 2; i < argc; i++){
				f << argv[i];
				cout << argv[i];
			}
			f << endl;
			cout << endl;
			f.close();
		}
		else {
			cout << "unable to open tasks.txt";
		}
	}
}
