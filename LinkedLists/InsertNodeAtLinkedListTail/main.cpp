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

        SinglyLinkedList() {
            this->head = nullptr;
        }

};

SinglyLinkedListNode* insertNodeAtTail(SinglyLinkedListNode* head, int data) {
    if (head == NULL){
        return head = new SinglyLinkedListNode(data);
    }
    SinglyLinkedListNode * curr = head;
    while(curr->next != NULL){
        curr = curr->next;
    }
    curr->next = new SinglyLinkedListNode(data);
    return head;
}

int main(){}