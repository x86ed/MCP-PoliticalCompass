package eightvalues

// Questions contains all the 8values political quiz questions
var Questions = []Question{
	{
		Index:  0,
		Text:   "Oppression by corporations is more of a concern than oppression by governments.",
		Effect: [4]float64{10, 0, -5, 0}, // econ, dipl, govt, scty
	},
	{
		Index:  1,
		Text:   "It is necessary for the government to intervene in the economy to protect consumers.",
		Effect: [4]float64{10, 0, 0, 0},
	},
	{
		Index:  2,
		Text:   "The freer the markets, the freer the people.",
		Effect: [4]float64{-10, 0, 0, 0},
	},
	{
		Index:  3,
		Text:   "It is better to maintain a balanced budget than to ensure welfare for all citizens.",
		Effect: [4]float64{-10, 0, 0, 0},
	},
	{
		Index:  4,
		Text:   "Publicly-funded research is more beneficial to the people than leaving it to the market.",
		Effect: [4]float64{10, 0, 0, 10},
	},
	{
		Index:  5,
		Text:   "Tariffs on international trade are important to encourage local production.",
		Effect: [4]float64{5, 0, -10, 0},
	},
	{
		Index:  6,
		Text:   "From each according to his ability, to each according to his needs.",
		Effect: [4]float64{10, 0, 0, 0},
	},
	{
		Index:  7,
		Text:   "It would be best if social programs were abolished in favor of private charity.",
		Effect: [4]float64{-10, 0, 0, 0},
	},
	{
		Index:  8,
		Text:   "Taxes should be increased on the rich to provide for the poor.",
		Effect: [4]float64{10, 0, 0, 0},
	},
	{
		Index:  9,
		Text:   "Inheritance is a legitimate form of wealth.",
		Effect: [4]float64{-10, 0, 0, -5},
	},
	{
		Index:  10,
		Text:   "Basic utilities like roads and electricity should be publicly owned.",
		Effect: [4]float64{10, 0, 0, 0},
	},
	{
		Index:  11,
		Text:   "Government intervention is a threat to the economy.",
		Effect: [4]float64{-10, 0, 0, 0},
	},
	{
		Index:  12,
		Text:   "Those with a greater ability to pay should receive better healthcare.",
		Effect: [4]float64{-10, 0, 0, 0},
	},
	{
		Index:  13,
		Text:   "Quality education is a right of all people.",
		Effect: [4]float64{10, 0, 0, 5},
	},
	{
		Index:  14,
		Text:   "The means of production should belong to the workers who use them.",
		Effect: [4]float64{10, 0, 0, 0},
	},
	{
		Index:  15,
		Text:   "The United Nations should be abolished.",
		Effect: [4]float64{0, -10, -5, 0},
	},
	{
		Index:  16,
		Text:   "Military action by our nation is often necessary to protect it.",
		Effect: [4]float64{0, -10, -10, 0},
	},
	{
		Index:  17,
		Text:   "I support regional unions, such as the European Union.",
		Effect: [4]float64{-5, 10, 10, 5},
	},
	{
		Index:  18,
		Text:   "It is important to maintain our national sovereignty.",
		Effect: [4]float64{0, -10, -5, 0},
	},
	{
		Index:  19,
		Text:   "A united world government would be beneficial to mankind.",
		Effect: [4]float64{0, 10, 0, 0},
	},
	{
		Index:  20,
		Text:   "It is more important to retain peaceful relations than to further our strength.",
		Effect: [4]float64{0, 10, 0, 0},
	},
	{
		Index:  21,
		Text:   "Wars do not need to be justified to other countries.",
		Effect: [4]float64{0, -10, -10, 0},
	},
	{
		Index:  22,
		Text:   "Military spending is a waste of money.",
		Effect: [4]float64{0, 10, 10, 0},
	},
	{
		Index:  23,
		Text:   "International aid is a waste of money.",
		Effect: [4]float64{-5, -10, 0, 0},
	},
	{
		Index:  24,
		Text:   "My nation is great.",
		Effect: [4]float64{0, -10, 0, 0},
	},
	{
		Index:  25,
		Text:   "Research should be conducted on an international scale.",
		Effect: [4]float64{0, 10, 0, 10},
	},
	{
		Index:  26,
		Text:   "Governments should be accountable to the international community.",
		Effect: [4]float64{0, 10, 5, 0},
	},
	{
		Index:  27,
		Text:   "Even when protesting an authoritarian government, violence is not acceptable.",
		Effect: [4]float64{0, 5, -5, 0},
	},
	{
		Index:  28,
		Text:   "My religious values should be spread as much as possible.",
		Effect: [4]float64{0, -5, -10, -10},
	},
	{
		Index:  29,
		Text:   "Our nation's values should be spread as much as possible.",
		Effect: [4]float64{0, -10, -5, 0},
	},
	{
		Index:  30,
		Text:   "It is very important to maintain law and order.",
		Effect: [4]float64{0, -5, -10, -5},
	},
	{
		Index:  31,
		Text:   "The general populace makes poor decisions.",
		Effect: [4]float64{0, 0, -10, 0},
	},
	{
		Index:  32,
		Text:   "Physician-assisted suicide should be legal.",
		Effect: [4]float64{0, 0, 10, 0},
	},
	{
		Index:  33,
		Text:   "The sacrifice of some civil liberties is necessary to protect us from acts of terrorism.",
		Effect: [4]float64{0, 0, -10, 0},
	},
	{
		Index:  34,
		Text:   "Government surveillance is necessary in the modern world.",
		Effect: [4]float64{0, 0, -10, 0},
	},
	{
		Index:  35,
		Text:   "The very existence of the state is a threat to our liberty.",
		Effect: [4]float64{0, 0, 10, 0},
	},
	{
		Index:  36,
		Text:   "Regardless of political opinions, it is important to side with your country.",
		Effect: [4]float64{0, -10, -10, -5},
	},
	{
		Index:  37,
		Text:   "All authority should be questioned.",
		Effect: [4]float64{0, 0, 10, 5},
	},
	{
		Index:  38,
		Text:   "A hierarchical state is best.",
		Effect: [4]float64{0, 0, -10, 0},
	},
	{
		Index:  39,
		Text:   "It is important that the government follows the majority opinion, even if it is wrong.",
		Effect: [4]float64{0, 0, 10, 0},
	},
	{
		Index:  40,
		Text:   "The stronger the leadership, the better.",
		Effect: [4]float64{0, -10, -10, 0},
	},
	{
		Index:  41,
		Text:   "Democracy is more than a decision-making process.",
		Effect: [4]float64{0, 0, 10, 0},
	},
	{
		Index:  42,
		Text:   "Environmental regulations are essential.",
		Effect: [4]float64{5, 0, 0, 10},
	},
	{
		Index:  43,
		Text:   "A better world will come from automation, science, and technology.",
		Effect: [4]float64{0, 0, 0, 10},
	},
	{
		Index:  44,
		Text:   "Children should be educated in religious or traditional values.",
		Effect: [4]float64{0, 0, -5, -10},
	},
	{
		Index:  45,
		Text:   "Traditions are of no value on their own.",
		Effect: [4]float64{0, 0, 0, 10},
	},
	{
		Index:  46,
		Text:   "Religion should play a role in government.",
		Effect: [4]float64{0, 0, -10, -10},
	},
	{
		Index:  47,
		Text:   "Churches should be taxed the same way other institutions are taxed.",
		Effect: [4]float64{5, 0, 0, 10},
	},
	{
		Index:  48,
		Text:   "Climate change is currently one of the greatest threats to our way of life.",
		Effect: [4]float64{0, 0, 0, 10},
	},
	{
		Index:  49,
		Text:   "It is important that we work as a united world to combat climate change.",
		Effect: [4]float64{0, 10, 0, 10},
	},
	{
		Index:  50,
		Text:   "Society was better many years ago than it is now.",
		Effect: [4]float64{0, 0, 0, -10},
	},
	{
		Index:  51,
		Text:   "It is important that we maintain the traditions of our past.",
		Effect: [4]float64{0, 0, 0, -10},
	},
	{
		Index:  52,
		Text:   "It is important that we think in the long term, beyond our lifespans.",
		Effect: [4]float64{0, 0, 0, 10},
	},
	{
		Index:  53,
		Text:   "Reason is more important than maintaining our culture.",
		Effect: [4]float64{0, 0, 0, 10},
	},
	{
		Index:  54,
		Text:   "Drug use should be legalized or decriminalized.",
		Effect: [4]float64{0, 0, 10, 2},
	},
	{
		Index:  55,
		Text:   "Same-sex marriage should be legal.",
		Effect: [4]float64{0, 0, 10, 10},
	},
	{
		Index:  56,
		Text:   "No cultures are superior to others.",
		Effect: [4]float64{0, 10, 5, 10},
	},
	{
		Index:  57,
		Text:   "Sex outside marriage is immoral.",
		Effect: [4]float64{0, 0, -5, -10},
	},
	{
		Index:  58,
		Text:   "If we accept migrants at all, it is important that they assimilate into our culture.",
		Effect: [4]float64{0, 0, -5, -10},
	},
	{
		Index:  59,
		Text:   "Abortion should be prohibited in most or all cases.",
		Effect: [4]float64{0, 0, -10, -10},
	},
	{
		Index:  60,
		Text:   "Gun ownership should be prohibited for those without a valid reason.",
		Effect: [4]float64{0, 0, -10, 0},
	},
	{
		Index:  61,
		Text:   "I support single-payer, universal healthcare.",
		Effect: [4]float64{10, 0, 0, 0},
	},
	{
		Index:  62,
		Text:   "Prostitution should be illegal.",
		Effect: [4]float64{0, 0, -10, -10},
	},
	{
		Index:  63,
		Text:   "Maintaining family values is essential.",
		Effect: [4]float64{0, 0, 0, -10},
	},
	{
		Index:  64,
		Text:   "To chase progress at all costs is dangerous.",
		Effect: [4]float64{0, 0, 0, -10},
	},
	{
		Index:  65,
		Text:   "Genetic modification is a force for good, even on humans.",
		Effect: [4]float64{0, 0, 0, 10},
	},
	{
		Index:  66,
		Text:   "We should open our borders to immigration.",
		Effect: [4]float64{0, 10, 10, 0},
	},
	{
		Index:  67,
		Text:   "Governments should be as concerned about foreigners as they are about their own citizens.",
		Effect: [4]float64{0, 10, 0, 0},
	},
	{
		Index:  68,
		Text:   "All people - regardless of factors like culture or sexuality - should be treated equally.",
		Effect: [4]float64{10, 10, 10, 10},
	},
	{
		Index:  69,
		Text:   "It is important that we further my group's goals above all others.",
		Effect: [4]float64{-10, -10, -10, -10},
	},
}
