# Domain Name System (DNS)

The Domain Name System is a hierarchical decentralized naming system for computers, services, or other resources connected to the Internet or a private network. It associates various information with domain names assigned to each of the participating entities. [Wikipedia](https://en.wikipedia.org/wiki/Domain_Name_System)

On a unix machine you can do a DNS lookup by running the command

    > host 190.74.13.200

    200.13.74.190.in-addr.arpa domain name pointer 190.74-13-200.dyn.dsl.cantv.net


    > host www.highiso.net
    www.highiso.net has address 109.74.193.253


    > host -t a cnn.com
    cnn.com has address 151.101.1.67
    cnn.com has address 151.101.129.67
    cnn.com has address 151.101.65.67
    cnn.com has address 151.101.193.67

