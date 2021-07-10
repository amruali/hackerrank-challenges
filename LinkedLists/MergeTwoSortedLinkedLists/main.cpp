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

SinglyLinkedListNode* mergeLists(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) {
  SinglyLinkedListNode * head = NULL;  
  SinglyLinkedListNode* Node = new SinglyLinkedListNode(-1);
  while(head1  || head2 )
  {  
      if (head1 && head2){
          if (head1 -> data <= head2 ->data){
                Node->next = head1;
                head1 = head1 ->next;
            }else{
             Node->next = head2;
             head2 = head2->next;
          }
      }
     else if (head1 != NULL){
          Node->next = head1;
          head1 = head1 ->next;
      }
      else if (head2 != NULL){
          Node->next = head2;
          head2 = head2->next;
      }else ;
      Node = Node->next;
      if (head == NULL){
          head = Node;
      }     
}
  return head;
}

int main()
{
    return 0;
}
