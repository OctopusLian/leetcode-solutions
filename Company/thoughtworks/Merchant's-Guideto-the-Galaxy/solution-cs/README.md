# Solution  

## CSharp  

### Program.cs  

```cs
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace RomanNumber
{
    class Program
    {

        static List<Roman> _romanList = new List<Roman>();
        static void Main(string[] args)
        {
            string glob = "I", prok = "V", pish = "X", tegj = "L";

            Console.WriteLine("glob is I");
            Console.WriteLine("prok is V");
            Console.WriteLine("pish is X");
            Console.WriteLine("tegj is L");
            Console.WriteLine("glob glob Silver is 34 Credits");
            Console.WriteLine("glob prok Gold is 57800 Credits");
            Console.WriteLine("pish pish Iron is 3910 Credits");
            Console.WriteLine("how much is pish tegj glob glob ?");
            Console.WriteLine("how many Credits is glob prok Silver ?");
            Console.WriteLine("how many Credits is glob prok Gold ?");
            Console.WriteLine("how many Credits is glob prok Iron ?");
            Console.WriteLine("how much wood could a woodchuck chuck if a woodchuck could chuck wood ?");
            Console.WriteLine("");
            Console.WriteLine("I was thinking about,please wait...");
            Console.WriteLine("");
            var romanService = RomanService.Instance;
            decimal Silver = 34m / romanService.Calculate(glob + glob);
            decimal Gold = 57800m / romanService.Calculate(glob + prok);
            decimal Iron = 3910m / romanService.Calculate(pish + pish);
            Console.WriteLine("pish tegj glob glob is {0}", romanService.Calculate(pish + tegj + glob + glob));
            Console.WriteLine("glob prok Silver is {0} Credits", romanService.Calculate(glob + prok) * Silver);
            Console.WriteLine("glob prok Gold is {0} Credits", romanService.Calculate(glob + prok) * Gold);
            Console.WriteLine("glob prok Iron is {0} Credits", romanService.Calculate(glob + prok) * Iron);
            Console.WriteLine("I have no idea what you are talking about");
            Console.WriteLine("please press any key to continue...");
            Console.ReadKey();
        }
    }

}
```

### Roman.cs  

```cs
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace RomanNumber
{
    /// <summary>
    /// A Featured Roman char 
    /// </summary>
    public class Roman
    {
        /// <summary>
        ///  I,V,X,L,C,D,M
        /// </summary>
        public char Symbol
        {
            set;
            get;
        }
        /// <summary>
        /// I:1,X:5,X:10,L:50,C:100,D:500,M:1000
        /// </summary>
        public int Value
        {
            set;
            get;
        }
        /// <summary>
        /// is allow repeated? such as II,III 
        /// only I,X,C,M can repeated
        /// </summary>
        public bool AllowRepeated
        {
            set;
            get;
        }
        /// <summary>
        /// is allow subtracted?
        /// only I,X C can subtracted
        /// such as: IV,IX; XL,XC;CD,CM;
        /// </summary>
        public bool AllowSubtracted
        {
            set;
            get;
        }
        /// <summary>
        /// if allow subtracted, how distence?
        /// I,V,X,L,C,D,M
        /// such as: 
        /// IV,IX:distence is 2; 
        /// XL,XC:distence is 2;
        /// CD,CM:distence is 2;
        /// </summary>
        public int SubtractedDestence
        {
            set;
            get;
        }
    }
}
```

### RomanService.cs  


```cs
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace RomanNumber
{
    public class RomanService
    {
        private char[] _allowRomans = new char[] { 'I', 'V', 'X', 'L', 'C', 'D', 'M' };

        /// <summary>
        /// The featured roman list
        /// </summary>
        public List<Roman> FeaturedRomans
        {
            private set;
            get;
        }

        private static RomanService _instance;
        public static RomanService Instance
        {
            get
            {
                if (_instance == null)
                    _instance = new RomanService();
                return _instance;
            }
        }

        private RomanService()
        {
            var vs = new int[] { 1, 5, 10, 50, 100, 500, 1000 };
            FeaturedRomans = new List<Roman>();
            for (var i = 0; i < _allowRomans.Length; i++)
            {
                FeaturedRomans.Add(new Roman
                {
                    Symbol = _allowRomans[i],
                    Value = vs[i],
                    AllowRepeated = _allowRomans[i] == 'I' || _allowRomans[i] == 'X' || _allowRomans[i] == 'C' || _allowRomans[i] == 'M',
                    AllowSubtracted = _allowRomans[i] == 'I' || _allowRomans[i] == 'X' || _allowRomans[i] == 'C',
                    SubtractedDestence = (_allowRomans[i] == 'I' || _allowRomans[i] == 'X' || _allowRomans[i] == 'C') ? 2 : 0
                });
            }
        }
        /// <summary>
        /// calculate a roman expression's value
        /// </summary>
        /// <param name="exp"></param>
        /// <returns></returns>
        public int Calculate(string exp)
        {
            if (string.IsNullOrEmpty(exp)) return 0;
            var romans = exp.ToCharArray();
            var result = 0;
            if (romans.Length == 1)
            {
                return FeaturedRomans.FirstOrDefault(x => x.Symbol == romans[0]).Value;
            }
            for (int i = 0; i <= romans.Length - 1; i = i + 2)
            {
                var current = FeaturedRomans.FirstOrDefault(x => x.Symbol == romans[i]);
                if (i >= romans.Length - 1)
                {
                    result += current.Value;
                    break;
                }
                var next = FeaturedRomans.FirstOrDefault(x => x.Symbol == romans[i + 1]);
                if (current == null)
                {
                    throw new Exception(string.Format("unexpected roman char:{0}", romans[i]));
                }
                if (next == null)
                {
                    throw new Exception(string.Format("unexpected roman char:{0}", romans[i + 1]));
                } 
                if (current.Value > next.Value)
                {
                    result += current.Value;
                    i--;
                }
                else if (current.Value == next.Value)
                {
                    if (!current.AllowRepeated)
                    {
                        throw new Exception(string.Format("Not allow {0} repeate", current.Symbol));
                    }
                    int repeatedCount = 2;
                    for (int j = i + 2; j < romans.Length; j++)
                    {
                        if (romans[j] != current.Symbol) break;
                        repeatedCount++;
                        if (repeatedCount > 3)
                        {
                            throw new Exception(string.Format("Not allow {0} repeated more than 3 times", current.Symbol));
                        }
                    }
                    result += current.Value;
                    i--;
                }
                else
                {
                    result += next.Value - current.Value;
                }
            }
            return result;
        }
    }
}
```