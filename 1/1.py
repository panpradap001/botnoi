def star(x):
    for i in range(1, x + 1):
        print("*" * i)

    for i in range(x-1,0,-1) : 
        print("*" * i)
        


x = int(input("Enter star : "))
star(x)