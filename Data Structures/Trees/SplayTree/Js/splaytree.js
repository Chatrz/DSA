function Node(key,val) 
{
	this.key = key;
	this.val = val;
	this.left = null;
	this.right = null;
}

function SplayBst() 
{
	this.root = null;
}

SplayBst.prototype.search = function(k) 
{
	if (this.root === null || ( !(Number(k) || k === 0) && typeof k !== "string"))
	return null;

	this.splay(k);
	return this.root.key === k ? this.root : null;
};

SplayBst.prototype.insert = function(k,v) 
{
	var n;
	if (( !(Number(k) || k === 0) && typeof k !== "string") || ( !(Number(v) || v === 0) && typeof v !== "string")) 
	{
		throw new Error("Invalid insert");
		return;
	}

	if (this.root === null) 
	{
		this.root = new Node(k,v);
		return;
	}

	this.splay(k);

	if (this.root.key > k) 
	{
		n = new Node(k,v);
		n.left = this.root.left;
		n.right = this.root;
		this.root.left = null;
		this.root = n;
	} else if (this.root.key < k) 
	{
		n = new Node(k,v);
		n.right = this.root.right;
		n.left = this.root;
		this.root.right = null;
		this.root = n;
	} else 
	{
		this.root.val = v;
	}
};

SplayBst.prototype.remove = function(k) 
{
	var temp;
	if (this.root === null || (!(Number(k) || k === 0) && typeof k !== "string"))
		return;

	this.splay(k);

	if (this.root.key === k) 
	{
		if (this.root.left === null && this.root.right === null) 
		{
		  this.root = null;
		} else if (this.root.left === null) 
		{
		  this.root = this.root.right;
		} else 
		{
		  temp = this.root.right;
		  this.root = this.root.left;
		  this.splay(k);
		  this.root.right = temp;
		}
	}
};

SplayBst.prototype.min = function(n) 
{
	var current;
	var minRecursive = function(cNode) 
	{
		if (cNode.left) 
		{
			return minRecursive(cNode.left);
		}
		return cNode;
	};

	if (this.root === null)
		return null;

	if (n instanceof Node)
		current = n;
	else
		current = this.root;

	return minRecursive(current);
};

SplayBst.prototype.max = function(n) 
{
	var current;
	var maxRecursive = function(cNode) 
	{
		if (cNode.right) 
		{
		  return maxRecursive(cNode.right);
		}
		return cNode;
	};

	if (this.root === null)
		return null;

	if (n instanceof Node)
		current = n;
	else
		current = this.root;

	return maxRecursive(current);
};

SplayBst.prototype.inOrder = function(n,fun) 
{
	if (n instanceof Node) 
	{
		this.inOrder(n.left,fun);
		if (fun) { fun(n); }
		this.inOrder(n.right,fun);
	}
};

SplayBst.prototype.contains = function(k) 
{
	var containsRecursive = function(n) 
	{
		if (n instanceof Node) 
		{
			if (n.key === k) 
			{
				return true;
			}
			containsRecursive(n.left);
			containsRecursive(n.right);
		}
	};

	if (this.root === null || (!(Number(k) || k === 0) && typeof k !== "string"))
		return false;

	return containsRecursive(this.root) ? true : false;
};

SplayBst.prototype.rotateRight = function(n)
{
	var temp;
	if (n instanceof Node) 
	{
		temp = n.left;
		n.left = temp.right;
		temp.right = n;
	}
	return temp;
};

SplayBst.prototype.rotateLeft = function(n) 
{
	var temp;
	if (n instanceof Node) 
	{
		temp = n.right;
		n.right = temp.left;
		temp.left = n;
	}
	return temp;
};

SplayBst.prototype.splay = function(k) 
{
	var splayRecursive = function(n, key) 
	{
		if (n === null)
		  return null;

		if (key < n.key) 
		{
			if (n.left === null) 
				return n;

			if (key < n.left.key) 
			{
				n.left.left = splayRecursive(n.left.left, key);
				n = this.rotateRight(n);
			} else if (key > n.left.key) 
			{
				n.left.right = splayRecursive(n.left.right, key);
				if (n.left.right !== null)
					n.left = this.rotateLeft(n.left);
			}
			if (n.left === null)
				return n;
			else 
				return this.rotateRight(n);
		} else if (key > n.key) 
		{
		  if (n.right === null) 
			return n;

		  if (key > n.right.key) 
		  {
			n.right.right = splayRecursive(n.right.right, key);
			n = this.rotateLeft(n);
		  } else if (key < n.right.key) 
		  {
			n.right.left = splayRecursive(n.right.left, key);
			if (n.right.left !== null)
			  n.right = this.rotateRight(n.right);
		  }
		  if (n.right === null)
			return n;
		  else 
			return this.rotateLeft(n);
		} else 
		{
		  return n;
		}

	}.bind(this);

	if (this.root === null || (!(Number(k) || k === 0) && typeof k !== "string")) 
	{
		throw new Error("Invalid splay");
		return;
	}

	this.root = splayRecursive(this.root,k);
	return;
};