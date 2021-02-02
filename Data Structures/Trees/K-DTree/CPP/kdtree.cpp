#include <algorithm>
#include <array>
#include <cmath>
#include <iostream>
#include <random>
#include <vector>
 

template<typename coordinate_type, size_t dimensions> class point 
{
    public:
        point(std::array<coordinate_type, dimensions> c) : coords_(c){}
        point(std::initializer_list<coordinate_type> list) 
        {
            size_t n = std::min(dimensions, list.size());
            std::copy_n(list.begin(), n, coords_.begin());
        }
        coordinate_type get(size_t index) const 
        {
            return coords_[index];
        }
        double distance(const point& pt) const 
        {
            double dist = 0;
            for (size_t i = 0; i < dimensions; ++i) 
            {
                double d = get(i) - pt.get(i);
                dist += d * d;
            }
            return dist;
        }
    private:
        std::array<coordinate_type, dimensions> coords_;
};
 

template<typename coordinate_type, size_t dimensions> std::ostream& operator<<(std::ostream& out, const point<coordinate_type, dimensions>& pt) 
{
    out << '(';
    for (size_t i = 0; i < dimensions; ++i) 
    {
        if (i > 0)
            out << ", ";
        out << pt.get(i);
    }
    out << ')';
    return out;
}
 

template<typename coordinate_type, size_t dimensions> class kdtree 
{
    public:
        typedef point<coordinate_type, dimensions> point_type;
    private:
        struct node 
        {
            node(const point_type& pt) : point_(pt), left_(nullptr), right_(nullptr){}
            coordinate_type get(size_t index) const 
            {
                return point_.get(index);
            }
            double distance(const point_type& pt) const 
            {
                return point_.distance(pt);
            }
            point_type point_;
            node* left_;
            node* right_;
        };
        node* root_ = nullptr;
        node* best_ = nullptr;
        double best_dist_ = 0;
        size_t visited_ = 0;
        std::vector<node> nodes_;
    
        struct node_cmp 
        {
            node_cmp(size_t index) : index_(index){}
            bool operator()(const node& n1, const node& n2) const 
            {
                return n1.point_.get(index_) < n2.point_.get(index_);
            }
            size_t index_;
        };
    
        node* make_tree(size_t begin, size_t end, size_t index) 
        {
            if (end <= begin)
                return nullptr;
            size_t n = begin + (end - begin)/2;
            std::nth_element(&nodes_[begin], &nodes_[n], &nodes_[0] + end, node_cmp(index));
            index = (index + 1) % dimensions;
            nodes_[n].left_ = make_tree(begin, n, index);
            nodes_[n].right_ = make_tree(n + 1, end, index);
            return &nodes_[n];
        }
    
        void nearest(node* root, const point_type& point, size_t index) 
        {
            if (root == nullptr)
                return;
            ++visited_;
            double d = root->distance(point);
            if (best_ == nullptr || d < best_dist_) 
            {
                best_dist_ = d;
                best_ = root;
            }
            if (best_dist_ == 0)
                return;
            double dx = root->get(index) - point.get(index);
            index = (index + 1) % dimensions;
            nearest(dx > 0 ? root->left_ : root->right_, point, index);
            if (dx * dx >= best_dist_)
                return;
            nearest(dx > 0 ? root->right_ : root->left_, point, index);
        }
    public:
        kdtree(const kdtree&) = delete;
        kdtree& operator=(const kdtree&) = delete;

        template<typename iterator> kdtree(iterator begin, iterator end) : nodes_(begin, end) 
        {
            root_ = make_tree(0, nodes_.size(), 0);
        }

        template<typename func> kdtree(func&& f, size_t n) 
        {
            nodes_.reserve(n);
            for (size_t i = 0; i < n; ++i)
                nodes_.emplace_back(f());
            root_ = make_tree(0, nodes_.size(), 0);
        }

        bool empty() const { return nodes_.empty(); }
    
        size_t visited() const { return visited_; }

        double distance() const { return std::sqrt(best_dist_); }
    
        const point_type& nearest(const point_type& pt) 
        {
            if (root_ == nullptr)
                throw std::logic_error("tree is empty");
            best_ = nullptr;
            visited_ = 0;
            best_dist_ = 0;
            nearest(root_, pt, 0);
            return best_->point_;
        }
};
 
void test_wikipedia() 
{
    typedef point<int, 2> point2d;
    typedef kdtree<int, 2> tree2d;
 
    point2d points[] = { { 2, 3 }, { 5, 4 }, { 9, 6 }, { 4, 7 }, { 8, 1 }, { 7, 2 } };
 
    tree2d tree(std::begin(points), std::end(points));
    point2d n = tree.nearest({ 9, 2 });
 
    std::cout << "Wikipedia example data:\n";
    std::cout << "nearest point: " << n << '\n';
    std::cout << "distance: " << tree.distance() << '\n';
    std::cout << "nodes visited: " << tree.visited() << '\n';
}
 
typedef point<double, 3> point3d;
typedef kdtree<double, 3> tree3d;
 
struct random_point_generator 
{
    random_point_generator(double min, double max) : engine_(std::random_device()()), distribution_(min, max) {}
 
    point3d operator()() 
    {
        double x = distribution_(engine_);
        double y = distribution_(engine_);
        double z = distribution_(engine_);
        return point3d({x, y, z});
    }
 
    std::mt19937 engine_;
    std::uniform_real_distribution<double> distribution_;
};
 
void test_random(size_t count) 
{
    random_point_generator rpg(0, 1);
    tree3d tree(rpg, count);
    point3d pt(rpg());
    point3d n = tree.nearest(pt);
 
    std::cout << "Random data (" << count << " points):\n";
    std::cout << "point: " << pt << '\n';
    std::cout << "nearest point: " << n << '\n';
    std::cout << "distance: " << tree.distance() << '\n';
    std::cout << "nodes visited: " << tree.visited() << '\n';
}
 
int main() 
{
    try 
    {
        test_wikipedia();
        std::cout << '\n';
        test_random(1000);
        std::cout << '\n';
        test_random(1000000);
    } catch (const std::exception& e) 
    {
        std::cerr << e.what() << '\n';
    }
    return 0;
}