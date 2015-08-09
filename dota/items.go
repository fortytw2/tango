package dota

// Item is something held by a Hero, typically...
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// String makes Item a Stringer
func (i *Item) String() string {
	return i.Name
}

// Items is a list of all current DOTA 2 items
// unobtainable via an API, so hardcoded here.
var Items = []Item{
	Item{
		ID:   0,
		Name: "Empty",
	},
	Item{
		ID:   1,
		Name: "Blink Dagger",
	},
	Item{
		ID:   2,
		Name: "Blades of Attack",
	},
	Item{
		ID:   3,
		Name: "Broadsword",
	},
	Item{
		ID:   4,
		Name: "Chainmail",
	},
	Item{
		ID:   5,
		Name: "Claymore",
	},
	Item{
		ID:   6,
		Name: "Helm of Iron Will",
	},
	Item{
		ID:   7,
		Name: "Javelin",
	},
	Item{
		ID:   8,
		Name: "Mithril Hammer",
	},
	Item{
		ID:   9,
		Name: "Platemail",
	},
	Item{
		ID:   10,
		Name: "Quarterstaff",
	},
	Item{
		ID:   11,
		Name: "Quelling Blade",
	},
	Item{
		ID:   12,
		Name: "Ring of Protection",
	},
	Item{
		ID:   182,
		Name: "Stout Shield",
	},
	Item{
		ID:   13,
		Name: "Gauntlets of Strength",
	},
	Item{
		ID:   14,
		Name: "Slippers of Agility",
	},
	Item{
		ID:   15,
		Name: "Mantle of Intelligence",
	},
	Item{
		ID:   16,
		Name: "GG Branch",
	},
	Item{
		ID:   17,
		Name: "Belt of Strength",
	},
	Item{
		ID:   18,
		Name: "boots_of_elves",
	},
	Item{
		ID:   19,
		Name: "Robe of the Magi",
	},
	Item{
		ID:   20,
		Name: "Circlet of Intelligence",
	},
	Item{
		ID:   21,
		Name: "Ogre Club",
	},
	Item{
		ID:   22,
		Name: "Blade of Alacrity",
	},
	Item{
		ID:   23,
		Name: "Staff of Wizardry",
	},
	Item{
		ID:   24,
		Name: "Ultimate Orb",
	},
	Item{
		ID:   25,
		Name: "gloves",
	},
	Item{
		ID:   26,
		Name: "lifesteal",
	},
	Item{
		ID:   27,
		Name: "ring_of_regen",
	},
	Item{
		ID:   28,
		Name: "sobi_mask",
	},
	Item{
		ID:   29,
		Name: "boots",
	},
	Item{
		ID:   30,
		Name: "gem",
	},
	Item{
		ID:   31,
		Name: "cloak",
	},
	Item{
		ID:   32,
		Name: "talisman_of_evasion",
	},
	Item{
		ID:   33,
		Name: "cheese",
	},
	Item{
		ID:   34,
		Name: "magic_stick",
	},
	Item{
		ID:   35,
		Name: "recipe_magic_wand",
	},
	Item{
		ID:   36,
		Name: "magic_wand",
	},
	Item{
		ID:   37,
		Name: "ghost",
	},
	Item{
		ID:   38,
		Name: "clarity",
	},
	Item{
		ID:   39,
		Name: "flask",
	},
	Item{
		ID:   40,
		Name: "dust",
	},
	Item{
		ID:   41,
		Name: "bottle",
	},
	Item{
		ID:   42,
		Name: "ward_observer",
	},
	Item{
		ID:   43,
		Name: "ward_sentry",
	},
	Item{
		ID:   44,
		Name: "tango",
	},
	Item{
		ID:   45,
		Name: "courier",
	},
	Item{
		ID:   46,
		Name: "tpscroll",
	},
	Item{
		ID:   47,
		Name: "recipe_travel_boots",
	},
	Item{
		ID:   48,
		Name: "travel_boots",
	},
	Item{
		ID:   49,
		Name: "recipe_phase_boots",
	},
	Item{
		ID:   50,
		Name: "phase_boots",
	},
	Item{
		ID:   51,
		Name: "demon_edge",
	},
	Item{
		ID:   52,
		Name: "eagle",
	},
	Item{
		ID:   53,
		Name: "reaver",
	},
	Item{
		ID:   54,
		Name: "relic",
	},
	Item{
		ID:   55,
		Name: "hyperstone",
	},
	Item{
		ID:   56,
		Name: "ring_of_health",
	},
	Item{
		ID:   57,
		Name: "void_stone",
	},
	Item{
		ID:   58,
		Name: "mystic_staff",
	},
	Item{
		ID:   59,
		Name: "energy_booster",
	},
	Item{
		ID:   60,
		Name: "point_booster",
	},
	Item{
		ID:   61,
		Name: "vitality_booster",
	},
	Item{
		ID:   62,
		Name: "recipe_power_treads",
	},
	Item{
		ID:   63,
		Name: "power_treads",
	},
	Item{
		ID:   64,
		Name: "recipe_hand_of_midas",
	},
	Item{
		ID:   65,
		Name: "hand_of_midas",
	},
	Item{
		ID:   66,
		Name: "recipe_oblivion_staff",
	},
	Item{
		ID:   67,
		Name: "oblivion_staff",
	},
	Item{
		ID:   68,
		Name: "recipe_pers",
	},
	Item{
		ID:   69,
		Name: "pers",
	},
	Item{
		ID:   70,
		Name: "recipe_poor_mans_shield",
	},
	Item{
		ID:   71,
		Name: "poor_mans_shield",
	},
	Item{
		ID:   72,
		Name: "recipe_bracer",
	},
	Item{
		ID:   73,
		Name: "bracer",
	},
	Item{
		ID:   74,
		Name: "recipe_wraith_band",
	},
	Item{
		ID:   75,
		Name: "wraith_band",
	},
	Item{
		ID:   76,
		Name: "recipe_null_talisman",
	},
	Item{
		ID:   77,
		Name: "null_talisman",
	},
	Item{
		ID:   78,
		Name: "recipe_mekansm",
	},
	Item{
		ID:   79,
		Name: "mekansm",
	},
	Item{
		ID:   80,
		Name: "recipe_vladmir",
	},
	Item{
		ID:   81,
		Name: "vladmir",
	},
	Item{
		ID:   84,
		Name: "flying_courier",
	},
	Item{
		ID:   85,
		Name: "recipe_buckler",
	},
	Item{
		ID:   86,
		Name: "buckler",
	},
	Item{
		ID:   87,
		Name: "recipe_ring_of_basilius",
	},
	Item{
		ID:   88,
		Name: "ring_of_basilius",
	},
	Item{
		ID:   89,
		Name: "recipe_pipe",
	},
	Item{
		ID:   90,
		Name: "pipe",
	},
	Item{
		ID:   91,
		Name: "recipe_urn_of_shadows",
	},
	Item{
		ID:   92,
		Name: "urn_of_shadows",
	},
	Item{
		ID:   93,
		Name: "recipe_headdress",
	},
	Item{
		ID:   94,
		Name: "headdress",
	},
	Item{
		ID:   95,
		Name: "recipe_sheepstick",
	},
	Item{
		ID:   96,
		Name: "sheepstick",
	},
	Item{
		ID:   97,
		Name: "recipe_orchid",
	},
	Item{
		ID:   98,
		Name: "orchid",
	},
	Item{
		ID:   99,
		Name: "recipe_cyclone",
	},
	Item{
		ID:   100,
		Name: "cyclone",
	},
	Item{
		ID:   101,
		Name: "recipe_force_staff",
	},
	Item{
		ID:   102,
		Name: "force_staff",
	},
	Item{
		ID:   103,
		Name: "recipe_dagon",
	},
	Item{
		ID:   197,
		Name: "recipe_dagon_2",
	},
	Item{
		ID:   198,
		Name: "recipe_dagon_3",
	},
	Item{
		ID:   199,
		Name: "recipe_dagon_4",
	},
	Item{
		ID:   200,
		Name: "recipe_dagon_5",
	},
	Item{
		ID:   104,
		Name: "dagon",
	},
	Item{
		ID:   201,
		Name: "dagon_2",
	},
	Item{
		ID:   202,
		Name: "dagon_3",
	},
	Item{
		ID:   203,
		Name: "dagon_4",
	},
	Item{
		ID:   204,
		Name: "dagon_5",
	},
	Item{
		ID:   105,
		Name: "recipe_necronomicon",
	},
	Item{
		ID:   191,
		Name: "recipe_necronomicon_2",
	},
	Item{
		ID:   192,
		Name: "recipe_necronomicon_3",
	},
	Item{
		ID:   106,
		Name: "necronomicon",
	},
	Item{
		ID:   193,
		Name: "necronomicon_2",
	},
	Item{
		ID:   194,
		Name: "necronomicon_3",
	},
	Item{
		ID:   107,
		Name: "recipe_ultimate_scepter",
	},
	Item{
		ID:   108,
		Name: "ultimate_scepter",
	},
	Item{
		ID:   109,
		Name: "recipe_refresher",
	},
	Item{
		ID:   110,
		Name: "refresher",
	},
	Item{
		ID:   111,
		Name: "recipe_assault",
	},
	Item{
		ID:   112,
		Name: "assault",
	},
	Item{
		ID:   113,
		Name: "recipe_heart",
	},
	Item{
		ID:   114,
		Name: "heart",
	},
	Item{
		ID:   115,
		Name: "recipe_black_king_bar",
	},
	Item{
		ID:   116,
		Name: "black_king_bar",
	},
	Item{
		ID:   117,
		Name: "aegis",
	},
	Item{
		ID:   118,
		Name: "recipe_shivas_guard",
	},
	Item{
		ID:   119,
		Name: "shivas_guard",
	},
	Item{
		ID:   120,
		Name: "recipe_bloodstone",
	},
	Item{
		ID:   121,
		Name: "bloodstone",
	},
	Item{
		ID:   122,
		Name: "recipe_sphere",
	},
	Item{
		ID:   123,
		Name: "sphere",
	},
	Item{
		ID:   124,
		Name: "recipe_vanguard",
	},
	Item{
		ID:   125,
		Name: "vanguard",
	},
	Item{
		ID:   126,
		Name: "recipe_blade_mail",
	},
	Item{
		ID:   127,
		Name: "blade_mail",
	},
	Item{
		ID:   128,
		Name: "recipe_soul_booster",
	},
	Item{
		ID:   129,
		Name: "soul_booster",
	},
	Item{
		ID:   130,
		Name: "recipe_hood_of_defiance",
	},
	Item{
		ID:   131,
		Name: "hood_of_defiance",
	},
	Item{
		ID:   132,
		Name: "recipe_rapier",
	},
	Item{
		ID:   133,
		Name: "rapier",
	},
	Item{
		ID:   134,
		Name: "recipe_monkey_king_bar",
	},
	Item{
		ID:   135,
		Name: "monkey_king_bar",
	},
	Item{
		ID:   136,
		Name: "recipe_radiance",
	},
	Item{
		ID:   137,
		Name: "radiance",
	},
	Item{
		ID:   138,
		Name: "recipe_butterfly",
	},
	Item{
		ID:   139,
		Name: "butterfly",
	},
	Item{
		ID:   140,
		Name: "recipe_greater_crit",
	},
	Item{
		ID:   141,
		Name: "greater_crit",
	},
	Item{
		ID:   142,
		Name: "recipe_basher",
	},
	Item{
		ID:   143,
		Name: "basher",
	},
	Item{
		ID:   144,
		Name: "recipe_bfury",
	},
	Item{
		ID:   145,
		Name: "bfury",
	},
	Item{
		ID:   146,
		Name: "recipe_manta",
	},
	Item{
		ID:   147,
		Name: "manta",
	},
	Item{
		ID:   148,
		Name: "recipe_lesser_crit",
	},
	Item{
		ID:   149,
		Name: "lesser_crit",
	},
	Item{
		ID:   150,
		Name: "recipe_armlet",
	},
	Item{
		ID:   151,
		Name: "armlet",
	},
	Item{
		ID:   183,
		Name: "recipe_invis_sword",
	},
	Item{
		ID:   152,
		Name: "invis_sword",
	},
	Item{
		ID:   153,
		Name: "recipe_sange_and_yasha",
	},
	Item{
		ID:   154,
		Name: "sange_and_yasha",
	},
	Item{
		ID:   155,
		Name: "recipe_satanic",
	},
	Item{
		ID:   156,
		Name: "satanic",
	},
	Item{
		ID:   157,
		Name: "recipe_mjollnir",
	},
	Item{
		ID:   158,
		Name: "mjollnir",
	},
	Item{
		ID:   159,
		Name: "recipe_skadi",
	},
	Item{
		ID:   160,
		Name: "skadi",
	},
	Item{
		ID:   161,
		Name: "recipe_sange",
	},
	Item{
		ID:   162,
		Name: "sange",
	},
	Item{
		ID:   163,
		Name: "recipe_helm_of_the_dominator",
	},
	Item{
		ID:   164,
		Name: "helm_of_the_dominator",
	},
	Item{
		ID:   165,
		Name: "recipe_maelstrom",
	},
	Item{
		ID:   166,
		Name: "maelstrom",
	},
	Item{
		ID:   167,
		Name: "recipe_desolator",
	},
	Item{
		ID:   168,
		Name: "desolator",
	},
	Item{
		ID:   169,
		Name: "recipe_yasha",
	},
	Item{
		ID:   170,
		Name: "yasha",
	},
	Item{
		ID:   171,
		Name: "recipe_mask_of_madness",
	},
	Item{
		ID:   172,
		Name: "mask_of_madness",
	},
	Item{
		ID:   173,
		Name: "recipe_diffusal_blade",
	},
	Item{
		ID:   195,
		Name: "recipe_diffusal_blade_2",
	},
	Item{
		ID:   174,
		Name: "diffusal_blade",
	},
	Item{
		ID:   196,
		Name: "diffusal_blade_2",
	},
	Item{
		ID:   175,
		Name: "recipe_ethereal_blade",
	},
	Item{
		ID:   176,
		Name: "ethereal_blade",
	},
	Item{
		ID:   177,
		Name: "recipe_soul_ring",
	},
	Item{
		ID:   178,
		Name: "soul_ring",
	},
	Item{
		ID:   179,
		Name: "recipe_arcane_boots",
	},
	Item{
		ID:   180,
		Name: "arcane_boots",
	},
	Item{
		ID:   181,
		Name: "orb_of_venom",
	},
	Item{
		ID:   184,
		Name: "recipe_ancient_janggo",
	},
	Item{
		ID:   185,
		Name: "ancient_janggo",
	},
	Item{
		ID:   186,
		Name: "recipe_medallion_of_courage",
	},
	Item{
		ID:   187,
		Name: "medallion_of_courage",
	},
	Item{
		ID:   188,
		Name: "smoke_of_deceit",
	},
	Item{
		ID:   189,
		Name: "recipe_veil_of_discord",
	},
	Item{
		ID:   190,
		Name: "veil_of_discord",
	},
	Item{
		ID:   205,
		Name: "recipe_rod_of_atos",
	},
	Item{
		ID:   206,
		Name: "rod_of_atos",
	},
	Item{
		ID:   207,
		Name: "recipe_abyssal_blade",
	},
	Item{
		ID:   208,
		Name: "abyssal_blade",
	},
	Item{
		ID:   209,
		Name: "recipe_heavens_halberd",
	},
	Item{
		ID:   210,
		Name: "heavens_halberd",
	},
	Item{
		ID:   211,
		Name: "recipe_ring_of_aquila",
	},
	Item{
		ID:   212,
		Name: "ring_of_aquila",
	},
	Item{
		ID:   213,
		Name: "recipe_tranquil_boots",
	},
	Item{
		ID:   214,
		Name: "tranquil_boots",
	},
	Item{
		ID:   215,
		Name: "shadow_amulet",
	},
	Item{
		ID:   216,
		Name: "enchanted_mango",
	},
	Item{
		ID:   218,
		Name: "ward_dispenser",
	},
	Item{
		ID:   220,
		Name: "travel_boots_2",
	},
	Item{
		ID:   226,
		Name: "lotus_orb",
	},
	Item{
		ID:   229,
		Name: "solar_crest",
	},
	Item{
		ID:   231,
		Name: "guardian_greaves",
	},
	Item{
		ID:   235,
		Name: "octarine_core",
	},
	Item{
		ID:   247,
		Name: "moon_shard",
	},
	Item{
		ID:   249,
		Name: "silver_edge",
	},
	Item{
		ID:   254,
		Name: "glimmer_cape",
	},
	Item{
		ID:   1000,
		Name: "halloween_candy_corn",
	},
	Item{
		ID:   1001,
		Name: "mystery_hook",
	},
	Item{
		ID:   1002,
		Name: "mystery_arrow",
	},
	Item{
		ID:   1003,
		Name: "mystery_missile",
	},
	Item{
		ID:   1004,
		Name: "mystery_toss",
	},
	Item{
		ID:   1005,
		Name: "mystery_vacuum",
	},
	Item{
		ID:   1006,
		Name: "halloween_rapier",
	},
	Item{
		ID:   1007,
		Name: "greevil_whistle",
	},
	Item{
		ID:   1008,
		Name: "greevil_whistle_toggle",
	},
	Item{
		ID:   1009,
		Name: "present",
	},
	Item{
		ID:   1010,
		Name: "winter_stocking",
	},
	Item{
		ID:   1011,
		Name: "winter_skates",
	},
	Item{
		ID:   1012,
		Name: "winter_cake",
	},
	Item{
		ID:   1013,
		Name: "winter_cookie",
	},
	Item{
		ID:   1014,
		Name: "winter_coco",
	},
	Item{
		ID:   1015,
		Name: "winter_ham",
	},
	Item{
		ID:   1016,
		Name: "winter_kringle",
	},
	Item{
		ID:   1017,
		Name: "winter_mushroom",
	},
	Item{
		ID:   1018,
		Name: "winter_greevil_treat",
	},
	Item{
		ID:   1019,
		Name: "winter_greevil_garbage",
	},
	Item{
		ID:   1020,
		Name: "winter_greevil_chewy",
	},
	Item{
		ID:   241,
		Name: "tango_single",
	},
	Item{
		ID:   242,
		Name: "crimson_guard",
	},
}
