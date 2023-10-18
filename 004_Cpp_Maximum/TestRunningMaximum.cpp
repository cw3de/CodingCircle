#include <gtest/gtest.h>
#include <iostream>
#include <list>
#include <random>
#include <vector>

std::vector<int> createRanomList( int size )
{
  std::mt19937 rng;
  rng.seed( std::random_device()() );
  std::uniform_int_distribution<std::mt19937::result_type> dist( 1, 100 );

  std::vector<int> result;
  for( int i = 0; i < size; ++i )
  {
    result.push_back( dist( rng ) );
  }
  return result;
}

std::ostream& operator<<( std::ostream& os, const std::vector<int>& v )
{
  os << "[";
  for( int i = 0; i < v.size(); ++i )
  {
    os << v[ i ];
    if( i < v.size() - 1 )
    {
      os << ", ";
    }
  }
  os << "]";
  return os;
}

// langsame variante, benötigt O(n*k) Zeit
std::vector<int> slowRunningMaximum( const std::vector<int>& V, int k )
{
  std::vector<int> result;
  for( int i = k; i <= V.size(); ++i )
  {
    int j = i - k;
    int max = V[ j ];
    for( ; j < i; ++j )
    {
      if( V[ j ] > max )
      {
        max = V[ j ];
      }
    }
    result.push_back( max );
  }
  return result;
}

// schnelle variante, benötigt O(n) Zeit
// Input: V= [ 27, 9, 17, 2, 12, 8 ], k = 3
// Output: [ 27, 17, 17, 12 ]

class FloatingMaximum
{
public:
  FloatingMaximum() = default;

  void addValue( int value )
  {
    // ersetzte alle kleineren Werte in der Queue, durch den neuen Wert
    int newCount{ 1 };
    int loopCount{ 0 };
    while( !maximaQueue.empty() && value >= maximaQueue.back().value )
    {
      newCount += maximaQueue.back().count;
      maximaQueue.pop_back();
      ++loopCount;
    }
    aufwand += loopCount ? loopCount : 1;
    maximaQueue.push_back( { value, newCount } );
  }

  int getNextMaximum()
  {
    const auto maxValue = maximaQueue.front();
    if( maxValue.count > 1 )
    {
      maximaQueue.front().count--;
    }
    else
    {
      maximaQueue.pop_front();
    }
    return maxValue.value;
  }

  int getAufwand() const { return aufwand; }

private:
  struct Maximum
  {
    int value;
    int count;
  };

  std::list<Maximum> maximaQueue; // constant time insert/remove at front/back
  int aufwand{ 0 };
};

std::vector<int> fastRunningMaximum( const std::vector<int>& V, int k )
{
  std::vector<int> result;
  FloatingMaximum fm;

  for( int i = 0; i < V.size(); ++i )
  {
    fm.addValue( V[ i ] );
    if( i >= k - 1 )
    {
      result.push_back( fm.getNextMaximum() );
    }
  }
  std::cout << "Aufwand für " << V.size() << " Elemente: " << fm.getAufwand() << std::endl;
  return result;
}

TEST( TestMaximum, EmptyList )
{
  std::vector<int> input = {};
  std::vector<int> expected = {};
  EXPECT_EQ( expected, slowRunningMaximum( input, 1 ) );
  EXPECT_EQ( expected, fastRunningMaximum( input, 1 ) );
}

TEST( TestMaximum, SmallList )
{
  std::vector<int> input = { 42, 27 };
  std::vector<int> expected = {};
  EXPECT_EQ( expected, slowRunningMaximum( input, 3 ) );
  EXPECT_EQ( expected, fastRunningMaximum( input, 3 ) );
}

TEST( TestMaximum, Test1List )
{
  std::vector<int> input = { 27, 9, 17, 2, 12, 8 };
  std::vector<int> expected = { 27, 17, 17, 12 };
  EXPECT_EQ( expected, slowRunningMaximum( input, 3 ) );
  EXPECT_EQ( expected, fastRunningMaximum( input, 3 ) );
}

TEST( TestMaximum, TestAscendingList )
{
  std::vector<int> input = { 1, 2, 3, 4, 5, 6, 7, 8, 9 };
  std::vector<int> expected = { 3, 4, 5, 6, 7, 8, 9 };
  EXPECT_EQ( expected, slowRunningMaximum( input, 3 ) );
  EXPECT_EQ( expected, fastRunningMaximum( input, 3 ) );
}

TEST( TestMaximum, TestDescendingList )
{
  std::vector<int> input = { 9, 8, 7, 6, 5, 4, 3, 2, 1 };
  std::vector<int> expected = { 9, 8, 7, 6, 5, 4, 3 };
  EXPECT_EQ( expected, slowRunningMaximum( input, 3 ) );
  EXPECT_EQ( expected, fastRunningMaximum( input, 3 ) );
}

TEST( TestMaximum, TestRampUpDownList )
{
  std::vector<int> input = { 9, 8, 7, 9, 8, 7, 9, 8, 7, 9, 8, 7, 9, 8, 7, 9, 8, 7 };
  std::vector<int> expected = slowRunningMaximum( input, 3 );
  EXPECT_EQ( expected, fastRunningMaximum( input, 3 ) );
}

TEST( TestMaximum, RandomList )
{
  const int size = 10000;
  const int k = 50;
  std::vector<int> input = createRanomList( size );
  // std::cout << "input:" << input << std::endl;
  std::vector<int> expected = slowRunningMaximum( input, k );
  EXPECT_EQ( input.size() - k + 1, expected.size() );
  // std::cout << "expected:    " << expected << std::endl;
  std::vector<int> computed = fastRunningMaximum( input, k );
  // std::cout << "computed:    " << computed << std::endl;
  EXPECT_EQ( expected, computed );
}

int main( int argc, char** argv )
{
  testing::InitGoogleTest( &argc, argv );
  return RUN_ALL_TESTS();
}