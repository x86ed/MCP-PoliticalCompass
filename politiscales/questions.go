package politiscales

// Questions contains all the politiscales questions with their weights
var Questions = []Question{
	{
		Index: 0,
		Text:  "constructivism_becoming_woman",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
			{Axis: "feminism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 1,
		Text:  "constructivism_racism_presence",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 2,
		Text:  "constructivism_science_society",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 3,
		Text:  "constructivism_gender_categories",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
			{Axis: "feminism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 4,
		Text:  "constructivism_criminality_nature",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 5,
		Text:  "constructivism_sexual_orientation",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 6,
		Text:  "constructivism_ethnic_differences",
		YesWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
	},
	{
		Index: 7,
		Text:  "essentialism_gender_biology",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
			{Axis: "feminism", Value: 3},
		},
	},
	{
		Index: 8,
		Text:  "essentialism_hormones_character",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
			{Axis: "feminism", Value: 3},
		},
	},
	{
		Index: 9,
		Text:  "essentialism_sexual_aggression",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
			{Axis: "feminism", Value: 3},
		},
	},
	{
		Index: 10,
		Text:  "essentialism_transgender_identity",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
	},
	{
		Index: 11,
		Text:  "essentialism_national_traits",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
	},
	{
		Index: 12,
		Text:  "essentialism_human_heterosexuality",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
	},
	{
		Index: 13,
		Text:  "essentialism_human_egoism",
		YesWeights: []Weight{
			{Axis: "essentialism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "constructivism", Value: 3},
		},
	},
	{
		Index: 14,
		Text:  "internationalism_border_removal",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 15,
		Text:  "internationalism_ideals_country",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 16,
		Text:  "internationalism_country_reparation",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 17,
		Text:  "internationalism_free_trade_similarity",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 18,
		Text:  "internationalism_sport_chauvinism",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 19,
		Text:  "internationalism_global_concern",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 20,
		Text:  "internationalism_foreign_political_rights",
		YesWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
	},
	{
		Index: 21,
		Text:  "nationalism_citizen_priority",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 22,
		Text:  "nationalism_country_values",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 23,
		Text:  "nationalism_multiculturalism_danger",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 24,
		Text:  "nationalism_good_citizen_patriot",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 25,
		Text:  "nationalism_military_intervention",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 26,
		Text:  "nationalism_history_national_belonging",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 27,
		Text:  "nationalism_country_research_access",
		YesWeights: []Weight{
			{Axis: "nationalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "internationalism", Value: 3},
		},
	},
	{
		Index: 28,
		Text:  "communism_wealth_ownership",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 29,
		Text:  "communism_private_labor_theft",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 30,
		Text:  "communism_public_health",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 31,
		Text:  "communism_public_energy_infrastructure",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 32,
		Text:  "communism_patents_nonexistence",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 33,
		Text:  "communism_production_rationing",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 34,
		Text:  "communism_labor_market_exploitation",
		YesWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
	},
	{
		Index: 35,
		Text:  "capitalism_profit_economy",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 36,
		Text:  "capitalism_merit_wealth_difference",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 37,
		Text:  "capitalism_private_schools_universities",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 38,
		Text:  "capitalism_relocation_production",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 39,
		Text:  "capitalism_rich_poor_acceptance",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 40,
		Text:  "capitalism_private_industry_sectors",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 41,
		Text:  "capitalism_private_banks",
		YesWeights: []Weight{
			{Axis: "capitalism", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "communism", Value: 3},
		},
	},
	{
		Index: 42,
		Text:  "regulation_income_tax_redistribution",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 43,
		Text:  "regulation_retirement_age",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 44,
		Text:  "regulation_unjustified_dismissals",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 45,
		Text:  "regulation_wage_control",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 46,
		Text:  "regulation_monopoly_prevention",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 47,
		Text:  "regulation_public_loans",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 48,
		Text:  "regulation_sector_subsidies",
		YesWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
	},
	{
		Index: 49,
		Text:  "laissez_faire_market_optimality",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 50,
		Text:  "laissez_faire_contract_freedom",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 51,
		Text:  "laissez_faire_labor_regulations",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 52,
		Text:  "laissez_faire_working_hours",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 53,
		Text:  "laissez_faire_environmental_standards",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 54,
		Text:  "laissez_faire_social_assistance",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 55,
		Text:  "laissez_faire_public_enterprises",
		YesWeights: []Weight{
			{Axis: "laissez_faire", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "regulation", Value: 3},
		},
	},
	{
		Index: 56,
		Text:  "progressive_tradition_questioning",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 57,
		Text:  "progressive_official_languages",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 58,
		Text:  "progressive_marriage_abolition",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 59,
		Text:  "progressive_foreign_culture_enrichment",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 60,
		Text:  "progressive_religion_influence",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 61,
		Text:  "progressive_language_definition",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 62,
		Text:  "progressive_euthanasia_legalization",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 63,
		Text:  "conservative_homosexual_equality",
		YesWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
	},
	{
		Index: 64,
		Text:  "conservative_death_penalty_justification",
		YesWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
	},
	{
		Index: 65,
		Text:  "conservative_technological_change",
		YesWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
	},
	{
		Index: 66,
		Text:  "conservative_school_curriculum",
		YesWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
	},
	{
		Index: 67,
		Text:  "conservative_abortion_restriction",
		YesWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
	},
	{
		Index: 68,
		Text:  "conservative_couple_child_production",
		YesWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
	},
	{
		Index: 69,
		Text:  "conservative_abstinence_preference",
		YesWeights: []Weight{
			{Axis: "conservative", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "progressive", Value: 3},
		},
	},
	{
		Index: 70,
		Text:  "ecology_species_extinction",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 71,
		Text:  "ecology_gmo_restriction",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 72,
		Text:  "ecology_climate_change_combat",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 73,
		Text:  "ecology_consumption_change",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 74,
		Text:  "ecology_biodiversity_agriculture",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 75,
		Text:  "ecology_ecosystem_preservation",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 76,
		Text:  "ecology_waste_reduction_production",
		YesWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "production", Value: 3},
		},
	},
	{
		Index: 77,
		Text:  "production_space_colonization",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 78,
		Text:  "production_ecosystem_transformation",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 79,
		Text:  "production_research_investment",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 80,
		Text:  "production_transhumanism_benefit",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 81,
		Text:  "production_nuclear_energy",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 82,
		Text:  "production_fossil_energy_exploitation",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 83,
		Text:  "production_economic_growth",
		YesWeights: []Weight{
			{Axis: "production", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "ecology", Value: 3},
		},
	},
	{
		Index: 84,
		Text:  "rehabilitative_justice_prison_abolition",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 85,
		Text:  "rehabilitative_justice_minimum_penalty",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 86,
		Text:  "rehabilitative_justice_reinsertion_support",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 87,
		Text:  "rehabilitative_justice_contextual_penalties",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 88,
		Text:  "rehabilitative_justice_detainee_conditions",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 89,
		Text:  "rehabilitative_justice_data_profiling",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 90,
		Text:  "rehabilitative_justice_internet_anonymity",
		YesWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
	},
	{
		Index: 91,
		Text:  "punitive_justice_punishment_goal",
		YesWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
	},
	{
		Index: 92,
		Text:  "punitive_justice_police_armed",
		YesWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
	},
	{
		Index: 93,
		Text:  "punitive_justice_terrorism_protection",
		YesWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
	},
	{
		Index: 94,
		Text:  "punitive_justice_order_authority",
		YesWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
	},
	{
		Index: 95,
		Text:  "punitive_justice_heavy_penalties_efficacy",
		YesWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
	},
	{
		Index: 96,
		Text:  "punitive_justice_preventive_arrest",
		YesWeights: []Weight{
			{Axis: "punitive_justice", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "rehabilitative_justice", Value: 3},
		},
	},
	{
		Index: 97,
		Text:  "revolution_general_strike_rights",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 98,
		Text:  "revolution_armed_struggle_necessity",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 99,
		Text:  "revolution_insurrection_necessity",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 100,
		Text:  "revolution_political_institutions",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 101,
		Text:  "revolution_election_challenge",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 102,
		Text:  "revolution_hacktivism_political",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 103,
		Text:  "revolution_sabotage_legitimacy",
		YesWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
	},
	{
		Index: 104,
		Text:  "reform_lawful_militation",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 105,
		Text:  "reform_revolution_outcome",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 106,
		Text:  "reform_radical_change_impact",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 107,
		Text:  "reform_violence_solution",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 108,
		Text:  "reform_manifestant_violence",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 109,
		Text:  "reform_opposition_compromise",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 110,
		Text:  "reform_individual_lifestyle_change",
		YesWeights: []Weight{
			{Axis: "reform", Value: 3},
		},
		NoWeights: []Weight{
			{Axis: "revolution", Value: 3},
		},
	},
	{
		Index: 111,
		Text:  "religion_diffusion",
		YesWeights: []Weight{
			{Axis: "religion", Value: 3},
		},
		NoWeights: []Weight{},
	},
	{
		Index: 112,
		Text:  "complotism_secret_control",
		YesWeights: []Weight{
			{Axis: "complotism", Value: 3},
		},
		NoWeights: []Weight{},
	},
	{
		Index: 113,
		Text:  "pragmatism_policy_approach",
		YesWeights: []Weight{
			{Axis: "pragmatism", Value: 3},
		},
		NoWeights: []Weight{},
	},
	{
		Index: 114,
		Text:  "monarchism_peace_sovereignty",
		YesWeights: []Weight{
			{Axis: "monarchism", Value: 3},
		},
		NoWeights: []Weight{},
	},
	{
		Index: 115,
		Text:  "veganism_animal_exploitation",
		YesWeights: []Weight{
			{Axis: "veganism", Value: 3},
		},
		NoWeights: []Weight{},
	},
	{
		Index: 116,
		Text:  "anarchism_state_abolition",
		YesWeights: []Weight{
			{Axis: "anarchism", Value: 3},
		},
		NoWeights: []Weight{},
	},
}
