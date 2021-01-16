def h1(key):
	x = (key + 7) * (key + 7)
	x = int(x / 16)
	x = x + key
	x = x % 11
	return x

def h(k , i):
	x = int( (i * i + i) / 2 )
	return ( h1(k) + x ) % 11


while True:
	order1 = int(input("<< "))
	order2 = int(input("<< "))
	if order1 == -1:
		break
	else:
		print(f'>> k: {order1} i:{order2} => {h(order1, order2)}')
