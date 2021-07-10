#include <bits/stdc++.h>

using namespace std;

class SinglyLinkedListNode {
    public:
        int data;
        SinglyLinkedListNode *next;

        SinglyLinkedListNode(int node_data) {
            this->data = node_data;
            this->next = nullptr;
        }
};

class SinglyLinkedList {
    public:
        SinglyLinkedListNode *head;
        SinglyLinkedListNode *tail;

        SinglyLinkedList() {
            this->head = nullptr;
            this->tail = nullptr;
        }

        void insert_node(int node_data) {
            SinglyLinkedListNode* node = new SinglyLinkedListNode(node_data);

            if (!this->head) {
                this->head = node;
            } else {
                this->tail->next = node;
            }

            this->tail = node;
        }
};


bool compare_lists(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) {
    SinglyLinkedListNode * currHead1 = head1;
    SinglyLinkedListNode * currHead2 = head2;
    
    
    while(currHead1 || currHead2){
        if (currHead1 && currHead2 && currHead1->data == currHead2->data)
        {
            currHead1 = currHead1->next;
            currHead2 = currHead2->next;
            continue;
        }
        return false;
    }
    return true;
}

int main(){
    /*
        Auto Genetated IO
    */
}