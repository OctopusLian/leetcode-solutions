# Solution  

## C++  

```cpp
#include <bits/stdc++.h>
#include <gtest/gtest.h>
 
/************* Common Define **************/
// error code definition // 定义错误码
typedef enum {
  eOK = 0, 
  eNullptr = -1, 
  eBadRequest = -2, 
  eNotValid = -3, 
  eCanceled = -4, 
  eBooked = -5, 
  eNotFound = -6, 
  eUnknownIssue = -999, 
}ErrCode;
 
// time format definition // 定义时间格式
typedef struct {
  int y, m, d;
  int hstart, hend;
  int w;
}Date;
 
// discount definition // 定义折扣信息
typedef struct {
  Date start;
  Date end;
  int discount;
}Discount;
 
// rename // 重命名类型
typedef std::vector<Discount> Discounts;
typedef std::string Place; // place: "A", "B", "C", "D"
typedef std::string UserName; // user name
typedef std::string CurrentOp; // current operation: "book", "cancel", "none"
typedef std::string LastOp; // last operation: "book", "cancel", "none"
typedef float BookPrice; // money paid for booking
typedef float CancelPrice; // money paid for canceling booking
typedef std::string ShortRecord; // string stored record in thr format of "yyyy-mm-dd-HH-HH-P-UserName"
 
// book record definition // 定义预定记录的格式
typedef struct {
  Date date;
  Place place;
  UserName user;
  CurrentOp cop;
  LastOp lop;
  BookPrice bp;
  CancelPrice cp;
}Record;
 
/* list to store the record in order // 此链表用来排序预定记录
析构函数未正确delete, 因此先注释掉
*/
typedef struct ShortRecordList {
  ShortRecordList(const ShortRecord &sr): short_record(sr), head(nullptr), tail(nullptr), next(nullptr) {}
  ~ShortRecordList(void) {
    ShortRecordList *c = this, *r = nullptr;
    while(nullptr != c) {
      r = c->next;
      // fprintf(stderr, "[DEBUG], %s > delete addr: %#x, %s\n", __func__, c, c->short_record.c_str());
      // delete c;
      c = r;
    }
  }
  // 添加一条记录, 并按照场馆使用时间升序排序
  ErrCode add(const ShortRecord &sr) {
    this->tail->next = new ShortRecordList(sr);
    return this->insertByOrder(this->tail->next);
  }
  // 排序
  ErrCode insertByOrder(ShortRecordList *srl) {
    auto c = this->tail->head;
    ShortRecordList *last = c;
    while(nullptr != c) {
      if(c->short_record < srl->short_record) {last = c; c = c->next;}
      else break;
    }
    if(this->tail->head == last) {
      // fprintf(stderr, "[DEBUG], %s > srl: %s -> last: %s, \n", __func__, srl->short_record.c_str(), last->short_record.c_str());
      this->tail->head->next = srl->next;
      srl->next = this->tail->head;
      this->tail->head = srl;
    }
    else {
      // fprintf(stderr, "[DEBUG], %s > last: %s -> srl: %s, \n", __func__, last->short_record.c_str(), srl->short_record.c_str());
      last->next = srl;
      srl->head = this->tail->head;
      this->tail = srl;
    }
    if(0)
    {
      c = this->tail->head;
      fprintf(stderr, "start -> ");
      while(nullptr != c) {fprintf(stderr, "%s -> ", c->short_record.c_str()); c = c->next;}
      fprintf(stderr, "end\n");
    }
    return eOK;
  }
  ShortRecord short_record; // 短记录, 上面已定义
  ShortRecordList *head; // 链表头指针, 调用格式为 plist->tail->head
  ShortRecordList *tail; // 链表尾指针, 调用格式为 plist->tail
  ShortRecordList *next; // 用来连接链表
}ShortRecordList;
 
typedef std::map<ShortRecord, Record> ShortRecordMap; // 映射短记录与完整记录
/************* Common Define **************/
 
/************* Booker **************/
/* 预定Booker类
shop_tradeable_matrix_: float二维数组, [行/列]为[星期几/几点钟], 存放每天每个时间点的价格
*/
class Booker {
public:
  Booker(void): output_(""), sum_(0.f), places_{"A", "B", "C", "D"}, 
                workday_cancel_percent_(0.5f), 
                weekday_cancel_percent_(workday_cancel_percent_ / 2.f) {
    for(auto &p: this->places_) {
      this->place_short_records_[p] = nullptr;
      this->place_price_sum_[p] = 0.f;
    }
    this->shop_tradeable_matrix_.resize(7);
    for(int w = 0; w < 7; ++w) {
      this->shop_tradeable_matrix_[w].resize(24, 0.f);
      for(int h = 0; h < 24; ++h) {
        if(5 == w or 6 == w) {
          if(h + 1 >=  9 and h + 1 < 12) this->shop_tradeable_matrix_[w][h] = 40.f;
          if(h + 1 >= 12 and h + 1 < 18) this->shop_tradeable_matrix_[w][h] = 50.f;
          if(h + 1 >= 18 and h + 1 < 22) this->shop_tradeable_matrix_[w][h] = 60.f;
        }
        else {
          if(h + 1 >=  9 and h + 1 < 12) this->shop_tradeable_matrix_[w][h] = 30.f;
          if(h + 1 >= 12 and h + 1 < 18) this->shop_tradeable_matrix_[w][h] = 50.f;
          if(h + 1 >= 18 and h + 1 < 20) this->shop_tradeable_matrix_[w][h] = 80.f;
          if(h + 1 >= 20 and h + 1 < 22) this->shop_tradeable_matrix_[w][h] = 60.f;
        }
      }
    }
    if(0)
    {
      fprintf(stderr, "shop working time matrix:\nday\\hour");
      for(int h = 7; h < 24; ++h) {
        fprintf(stderr, "%2d\t", h + 1);
      }
      fprintf(stderr, "\n");
      for(int w = 0; w < 7; ++w) {
        fprintf(stderr, "%d\t", w + 1);
        for(int h = 7; h < 24; ++h) {
          fprintf(stderr, "%2.0f\t", this->shop_tradeable_matrix_[w][h]);
        }
        fprintf(stderr, "\n");
      }
    }
  }
   
  ~Booker(void) {
    for(auto &psr: this->place_short_records_) {
      if(nullptr != psr.second) delete psr.second->head;
    }
  }
   
  // 获取输入并处理该条记录
  const std::string getInput(const std::string &inp) {
    auto status = eNullptr;
    Record record = {0};
    ShortRecord short_record = "";
    status = this->parseRecordFromString(inp, record, short_record);
    if(eOK != status) {
      return "Error: the booking is invalid!";
    }
    // 添加/修改记录
    // fprintf(stderr, "[DEBUG] >>> \n");
    if(0 == this->short_record_map_.count(short_record)) {
      if(nullptr == this->place_short_records_[record.place]) {
        this->place_short_records_[record.place] = new ShortRecordList(short_record);
        this->place_short_records_[record.place]->head = \
        this->place_short_records_[record.place]->tail = \
        this->place_short_records_[record.place];
      }
      else this->place_short_records_[record.place]->add(short_record);
      this->short_record_map_[short_record] = record;
      // fprintf(stderr, "[DEBUG], %s > insert %s, cop: %s, lop: %s\n", \
              __func__, short_record.c_str(), \
              this->short_record_map_[short_record].cop.c_str(), record.lop.c_str());
    }
    else {
      this->short_record_map_[short_record].cop = record.cop;
      // fprintf(stderr, "[DEBUG], %s > skip %s, cop: %s, lop: %s\n", \
              __func__, short_record.c_str(), \
              this->short_record_map_[short_record].cop.c_str(),  \
              this->short_record_map_[short_record].lop.c_str());
    }
    // 处理记录
    status = this->go(short_record);
    if(eCanceled == status) {
      return "Error: the booking being cancelled does not exist!";
    }
    if(eBooked == status) {
      return "Error: the booking conflicts with existing bookings!";
    }
    else if(eOK != status){
      return "Bad return in go!";
    }
    return "Success: the booking is accepted!";
  }
   
  // 分割字符串
  const std::vector<std::string> split(const std::string &line, const std::string c = " ") {
    std::vector<std::string> v;
    std::string::size_type pos1, pos2;
    pos2 = line.find(c);
    pos1 = 0;
    while(std::string::npos != pos2) {
      v.push_back(line.substr(pos1, pos2-pos1));       
      pos1 = pos2 + c.size();
      pos2 = line.find(c, pos1);
    }
    if(pos1 != line.length()) v.push_back(line.substr(pos1));
    return v;
  }
   
  // 解析字符串到记录
  ErrCode parseRecordFromString(const std::string &line, Record &record, ShortRecord &short_record) {
    auto v = this->split(line, " ");
    if(0)
    {
      fprintf(stderr, "size: %ld\n", v.size());
      for(auto &i: v) fprintf(stderr, "%s ", i.c_str());
      fprintf(stderr, "\n");
    }
    return this->parseAndValid(v, record, short_record);
  }
   
  // 解析并验证是否为有效输入
  ErrCode parseAndValid(const std::vector<std::string> &vr, Record &record, ShortRecord &short_record) {
    if(4 != vr.size() and 5 != vr.size()) return eNotValid;
    if(5 == vr.size()) {
      if(vr[4] != "C") return eNotValid;
      else record.cop = "cancel";
    }
    else record.cop = "book";
    // user yyyy-mm-dd HH:MM~HH:MM place [cop]
    record.user = vr[0];
    auto status = this->parseDate(vr[1], vr[2], record.date);
    if(eOK != status) return status;
    record.place = vr[3];
    if(vr[3] != "A" and vr[3] != "B" and vr[3] != "C" and vr[3] != "D") return eNotValid;
    record.lop = "none";
    short_record = vr[1] + "-" + vr[2].substr(0, 2) + "-" + vr[2].substr(6, 2) + "-" + vr[3] + "-" + vr[0];
    // this->print(record);
    return eOK;
  }
   
  // 解析时间
  ErrCode parseDate(const std::string &ymd, const std::string &hm, Date &date) {
    auto v = this->split(ymd, "-");
    if(3 != v.size()) return eNotValid;
    date.y = std::stoi(v[0]);
    date.m = std::stoi(v[1]);
    date.d = std::stoi(v[2]);
    v = this->split(hm, "~");
    if(2 != v.size()) return eNotValid;
    int m = -1;
    auto starthm = this->split(v[0], ":");
    auto endhm = this->split(v[1], ":");
    if(2 != starthm.size() or 2 != endhm.size()) return eNotValid;
    m = std::stoi(starthm[1] + endhm[1]);
    if(0 != m) return eNotValid;
    date.hstart = std::stoi(starthm[0]);
    date.hend = std::stoi(endhm[0]);
    return this->valid(date);
  }
     
  // 验证是否有效
  ErrCode valid(Date &date) {
    // evaluate yyyy-mm-dd, not implemented
    if(date.hstart >= date.hend) return eNotValid;
    return this->shopTradeable(date);
  }
   
  // 是否在营业时间内
  ErrCode shopTradeable(Date &date) {
    int w = date.w = this->getDayByDate(date);
     
    if(0 == this->shop_tradeable_matrix_[w - 1][date.hstart - 1] or \
       0 == this->shop_tradeable_matrix_[w - 1][date.hend - 2]) return eNotValid;
    return eOK;
  }
   
  // 根据时间计算星期几
  int getDayByDate(const Date &date) {
    return this->getDayByDate(date.y, date.m, date.d);
  }
   
  // 根据时间计算星期几
  int getDayByDate(int y, int m, int d) {
    if(m == 1 || m == 2) {m += 12; y--;}
    int w = 1 + (d + 2 * m + 3 * (m + 1) / 5 + y + y / 4 - y / 100 + y / 400 ) % 7;
    return w;
  }
   
  // 处理记录
  ErrCode go(const ShortRecord &short_record) {
    auto &record = this->short_record_map_[short_record];
    // fprintf(stderr, "[DEBUG], start %s > short_record: %s, cop: %s, lop: %s\n", __func__, \
            short_record.c_str(), record.cop.c_str(), record.lop.c_str());
    if(record.cop == "cancel") {
      return this->cancel(short_record);
    }
    else if(record.cop == "book") {
      return this->book(short_record);
    }
    return eBadRequest;
  }
   
  // 获取折扣信息, 此处未实现细节, 默认不打折
  BookPrice fetchDiscount(const Record &record) {
     
    return 1.f;
  }
   
  // 计算预订价格
  BookPrice calcBookPrice(const Record &record) {
    BookPrice price = 0.f;
    for(int i = record.date.hstart - 1; i < record.date.hend - 1; ++i) {
      price += (this->shop_tradeable_matrix_[record.date.w][i] * this->fetchDiscount(record));
    }
    return price;
  }
   
  // 是否为同一天同一地点
  bool isTheSameDayAndPlace(const ShortRecord &sr1, const ShortRecord &sr2) {
    return (sr1.substr(0, 10) == sr2.substr(0, 10) and sr1[17] == sr2[17]);
  }
   
  // 时间是否有重叠
  bool hasIntersection(const ShortRecord &sr1, const ShortRecord &sr2) {    
    auto &r1 = this->short_record_map_[sr1];
    auto &r2 = this->short_record_map_[sr2];
    if(sr1 == sr2) {
      if(r1.lop == "book") return true;
      return false;
    }
    int s1 = r1.date.hstart, e1 = r1.date.hend, s2 = r2.date.hstart, e2 = r2.date.hend;
    // fprintf(stderr, "[%2d~%2d] [%2d~%2d]\n", s1, e1, s2, e2);
    if((s1 >= s2 and s1 < e2) or (e1 > s2 and e1 <= e2) or (s2 >= s1 and s2 < e1) or (e2 > s1 and e2 <= e1)) if(r1.lop == "book") return true;
    return false;
  }
   
  // 是否被预定
  bool isBooked(const ShortRecord &short_record) {
    bool start = false;
    for(auto &sr: this->short_record_map_) {
      // fprintf(stderr, "[DEBUG], %s > short_record_map: %s\n", __func__, sr.first.c_str());
      if(this->isTheSameDayAndPlace(sr.first, short_record)) {start = true;} else {if(start) break; else continue;}
      if(this->hasIntersection(sr.first, short_record)) return true;
    }
    return false;
  }
   
  // 预定
  ErrCode book(const ShortRecord &short_record) {
    // fprintf(stderr, "[DEBUG] > booking...\n");
    if(this->isBooked(short_record)) return eBooked;
    auto &record = this->short_record_map_[short_record];
    record.lop = record.cop; record.cop = "none";
    record.bp = this->calcBookPrice(record);
    // fprintf(stderr, "[DEBUG], after %s > short_record: %s, cop: %s, lop: %s\n", __func__,  \
            short_record.c_str(), record.cop.c_str(), record.lop.c_str());
    return eOK;
  }
   
  // 取消
  ErrCode cancel(const ShortRecord &short_record) {
    // fprintf(stderr, "[DEBUG] > canceling...\n");
    int found = this->short_record_map_.count(short_record);
    if(0 == found) return eNotFound;
    auto &record = this->short_record_map_[short_record];
    if(record.lop == "cancel") return eCanceled;
    record.lop = record.cop; record.cop = "none";
    record.cp = this->calcBookPrice(record) * (record.date.w >= 6? this->weekday_cancel_percent_: this->workday_cancel_percent_);
    // fprintf(stderr, "[DEBUG], after %s > short_record: %s, cop: %s, lop: %s\n", __func__, \
            short_record.c_str(), record.cop.c_str(), record.lop.c_str());
    return eOK;
  }
   
  // 输出
  const std::string getOutput(void) {
    fprintf(stderr, "/************************* OUTPUT *************************/\n");
    int i = 0;
    fprintf(stderr, "> 收入汇总\n> ---");
    for(auto &psr: this->place_short_records_) {
      if(i++ != 0) {
        fprintf(stderr, ">");
      }
      fprintf(stderr, "\n> 场地:%s\n", psr.first.c_str());
      ShortRecordList* srl = nullptr;
      if(nullptr != psr.second) srl = psr.second->head;
      while(nullptr != srl) {
        auto &r = this->short_record_map_[srl->short_record];
        this->print(r, false);
        srl = srl->next;
      }
      // for(auto &psri: psr.second) {
        // auto &r = this->short_record_map_[psri];
        // this->print(r, false);
      // }
      fprintf(stderr, "> 小计: %.0f元\n", this->place_price_sum_[psr.first]);
    }
    fprintf(stderr, "> ---\n> 总计: %.0f元\n", this->sum_);
    return this->output_;
  }
   
  // 打印记录
  void print(const Record &record, bool full = true) {
    if(full) {
      fprintf(stderr, "[DEBUG] > %s %04d-%02d-%02d %02d:00~%02d:00 %s [%s]\n", \
              record.user.c_str(), record.date.y, record.date.m, record.date.d, \
              record.date.hstart, record.date.hend, record.place.c_str(), record.cop.c_str());
    }
    else {
      fprintf(stderr, "> %04d-%02d-%02d %02d:00~%02d:00 %s%.0f元\n", \
              record.date.y, record.date.m, record.date.d, \
              record.date.hstart, record.date.hend, \
              (record.lop == "cancel"? "违约金 ": ""), (record.lop == "cancel"? record.cp: record.bp));
      if(record.lop == "cancel") {
        this->sum_ += record.cp;
        this->place_price_sum_[record.place] += record.cp;
         
      }
      else {
        this->sum_ += record.bp;
        this->place_price_sum_[record.place] += record.bp;
      }
    }
  }
   
protected:
  std::vector<std::string> places_;
  std::map<Place, ShortRecordList*> place_short_records_;
  ShortRecordMap short_record_map_;
  std::map<Place, BookPrice> place_price_sum_;
  std::string output_;
  std::vector<std::vector<float> > shop_tradeable_matrix_;
  BookPrice sum_;
  CancelPrice workday_cancel_percent_;
  CancelPrice weekday_cancel_percent_;
};
/************* Booker **************/
 
/************* TEST **************/
TEST(BookerTest, test1) {
  std::vector<std::string> input_testinstance = {
    "abcdefghijklmnopqrst1234567890", 
    "U001 2016-06-02 22:00~22:00 A", 
    "U002 2017-08-01 19:00~22:00 A", 
    "U003 2017-08-02 13:00~17:00 B", 
    "U004 2017-08-03 15:00~16:00 C", 
    "U005 2017-08-05 09:00~11:00 D", 
  };
   
  std::vector<std::string> input_out_testinstance = {
    "Error: the booking is invalid!", 
    "Error: the booking is invalid!", 
    "Success: the booking is accepted!", 
    "Success: the booking is accepted!", 
    "Success: the booking is accepted!", 
    "Success: the booking is accepted!", 
  };
   
  const std::string output = "\
> 收入汇总\n\
> ---\n\
> 场地:A\n\
> 2017-08-01 19:00~22:00 200元\n\
> 小计: 200元\n\
>\n\
> 场地:B\n\
> 2017-08-02 13:00~17:00 200元\n\
> 小计: 200元\n\
>\n\
> 场地:C\n\
> 2017-08-03 15:00~16:00 50元\n\
> 小计: 50元\n\
>\n\
> 场地:D\n\
> 2017-08-05 09:00~11:00 80元\n\
> 小计: 80元\n\
> ---\n\
> 总计: 530元\n\
";
   
  auto p = new Booker;
  ASSERT_TRUE(p != nullptr);
  for(int i = 0; i < input_testinstance.size(); ++i) {
    EXPECT_EQ(p->getInput(input_testinstance[i]), input_out_testinstance[i]);
  }
  fprintf(stderr, "%s", output.c_str());
  p->getOutput();
  delete p;
}
 
TEST(BookerTest, test2) {
  std::vector<std::string> input_testinstance = {
    "U002 2017-08-01 19:00~22:00 A", 
    "U003 2017-08-01 18:00~20:00 A", 
    "U002 2017-08-01 19:00~22:00 A C", 
    "U002 2017-08-01 19:00~22:00 A C", 
    "U003 2017-08-01 18:00~20:00 A", 
    "U003 2017-08-02 13:00~17:00 B", 
  };
   
  std::vector<std::string> input_out_testinstance = {
    "Success: the booking is accepted!", 
    "Error: the booking conflicts with existing bookings!", 
    "Success: the booking is accepted!", 
    "Error: the booking being cancelled does not exist!", 
    "Success: the booking is accepted!", 
    "Success: the booking is accepted!", 
  };
   
  const std::string output = "\
> 收入汇总\n\
> ---\n\
> 场地:A\n\
> 2017-08-01 18:00~20:00 160元\n\
> 2017-08-01 19:00~22:00 违约金 100元\n\
> 小计：260元\n\
>\n\
> 场地:B\n\
> 2017-08-02 13:00~17:00 200元\n\
> 小计：200元\n\
>\n\
> 场地:C\n\
> 小计：0元\n\
>\n\
> 场地:D\n\
> 小计：0元\n\
> ---\n\
> 总计：460元\n\
";
   
  auto p = new Booker;
  ASSERT_TRUE(p != nullptr);
  for(int i = 0; i < input_testinstance.size(); ++i) {
    EXPECT_EQ(p->getInput(input_testinstance[i]), input_out_testinstance[i]);
  }
  fprintf(stderr, "%s", output.c_str());
  p->getOutput();
  delete p;
}
 
TEST(BookerTest, test3) {
  std::vector<std::string> input_testinstance = {
    "U002 2017-08-01 19:00~22:00 A", 
    "U003 2017-08-01 18:00~20:00 A", 
    "U002 2017-08-01 19:00~22:00 A C", 
    "U002 2017-08-01 19:00~22:00 A C", 
    "U003 2017-08-01 18:00~20:00 A", 
    "U003 2017-08-02 13:00~17:00 B", 
    "U003 2017-08-02 13:00~17:00 B D", 
    "U003 2017-08-02 13:00~17:00 B", 
  };
   
  std::vector<std::string> input_out_testinstance = {
    "Success: the booking is accepted!", 
    "Error: the booking conflicts with existing bookings!", 
    "Success: the booking is accepted!", 
    "Error: the booking being cancelled does not exist!", 
    "Success: the booking is accepted!", 
    "Success: the booking is accepted!", 
    "Error: the booking is invalid!", 
    "Error: the booking conflicts with existing bookings!", 
  };
   
  const std::string output = "\
> 收入汇总\n\
> ---\n\
> 场地:A\n\
> 2017-08-01 18:00~20:00 160元\n\
> 2017-08-01 19:00~22:00 违约金 100元\n\
> 小计：260元\n\
>\n\
> 场地:B\n\
> 2017-08-02 13:00~17:00 200元\n\
> 小计：200元\n\
>\n\
> 场地:C\n\
> 小计：0元\n\
>\n\
> 场地:D\n\
> 小计：0元\n\
> ---\n\
> 总计：460元\n\
";
   
  auto p = new Booker;
  ASSERT_TRUE(p != nullptr);
  for(int i = 0; i < input_testinstance.size(); ++i) {
    EXPECT_EQ(p->getInput(input_testinstance[i]), input_out_testinstance[i]);
  }
  fprintf(stderr, "%s", output.c_str());
  p->getOutput();
  delete p;
}
/************* TEST **************/
 
int main(int argc, char **argv) {
  testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();;
}
```