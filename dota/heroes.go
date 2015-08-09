package dota

// A Hero maps 1-to-1 with a DOTA 2 hero, providing all relevant information.
type Hero struct {
	ID            int    `json:"id"`
	Name          string `json:"-"`
	LocalizedName string `json:"name"`
}

// Heroes is a list of all Heroes currently in Dota 2
var Heroes = []Hero{
	Hero{
		Name:          "antimage",
		ID:            1,
		LocalizedName: "Anti-Mage",
	},
	Hero{
		Name:          "axe",
		ID:            2,
		LocalizedName: "Axe",
	},
	Hero{
		Name:          "bane",
		ID:            3,
		LocalizedName: "Bane",
	},
	Hero{
		Name:          "bloodseeker",
		ID:            4,
		LocalizedName: "Bloodseeker",
	},
	Hero{
		Name:          "crystal_maiden",
		ID:            5,
		LocalizedName: "Crystal Maiden",
	},
	Hero{
		Name:          "drow_ranger",
		ID:            6,
		LocalizedName: "Drow Ranger",
	},
	Hero{
		Name:          "earthshaker",
		ID:            7,
		LocalizedName: "Earthshaker",
	},
	Hero{
		Name:          "juggernaut",
		ID:            8,
		LocalizedName: "Juggernaut",
	},
	Hero{
		Name:          "mirana",
		ID:            9,
		LocalizedName: "Mirana",
	},
	Hero{
		Name:          "nevermore",
		ID:            11,
		LocalizedName: "Shadow Fiend",
	},
	Hero{
		Name:          "morphling",
		ID:            10,
		LocalizedName: "Morphling",
	},
	Hero{
		Name:          "phantom_lancer",
		ID:            12,
		LocalizedName: "Phantom Lancer",
	},
	Hero{
		Name:          "puck",
		ID:            13,
		LocalizedName: "Puck",
	},
	Hero{
		Name:          "pudge",
		ID:            14,
		LocalizedName: "Pudge",
	},
	Hero{
		Name:          "razor",
		ID:            15,
		LocalizedName: "Razor",
	},
	Hero{
		Name:          "sand_king",
		ID:            16,
		LocalizedName: "Sand King",
	},
	Hero{
		Name:          "storm_spirit",
		ID:            17,
		LocalizedName: "Storm Spirit",
	},
	Hero{
		Name:          "sven",
		ID:            18,
		LocalizedName: "Sven",
	},
	Hero{
		Name:          "tiny",
		ID:            19,
		LocalizedName: "Tiny",
	},
	Hero{
		Name:          "vengefulspirit",
		ID:            20,
		LocalizedName: "Vengeful Spirit",
	},
	Hero{
		Name:          "windrunner",
		ID:            21,
		LocalizedName: "Windranger",
	},
	Hero{
		Name:          "zuus",
		ID:            22,
		LocalizedName: "Zeus",
	},
	Hero{
		Name:          "kunkka",
		ID:            23,
		LocalizedName: "Kunkka",
	},
	Hero{
		Name:          "lina",
		ID:            25,
		LocalizedName: "Lina",
	},
	Hero{
		Name:          "lich",
		ID:            31,
		LocalizedName: "Lich",
	},
	Hero{
		Name:          "lion",
		ID:            26,
		LocalizedName: "Lion",
	},
	Hero{
		Name:          "shadow_shaman",
		ID:            27,
		LocalizedName: "Shadow Shaman",
	},
	Hero{
		Name:          "slardar",
		ID:            28,
		LocalizedName: "Slardar",
	},
	Hero{
		Name:          "tidehunter",
		ID:            29,
		LocalizedName: "Tidehunter",
	},
	Hero{
		Name:          "witch_doctor",
		ID:            30,
		LocalizedName: "Witch Doctor",
	},
	Hero{
		Name:          "riki",
		ID:            32,
		LocalizedName: "Riki",
	},
	Hero{
		Name:          "enigma",
		ID:            33,
		LocalizedName: "Enigma",
	},
	Hero{
		Name:          "tinker",
		ID:            34,
		LocalizedName: "Tinker",
	},
	Hero{
		Name:          "sniper",
		ID:            35,
		LocalizedName: "Sniper",
	},
	Hero{
		Name:          "necrolyte",
		ID:            36,
		LocalizedName: "Necrophos",
	},
	Hero{
		Name:          "warlock",
		ID:            37,
		LocalizedName: "Warlock",
	},
	Hero{
		Name:          "beastmaster",
		ID:            38,
		LocalizedName: "Beastmaster",
	},
	Hero{
		Name:          "queenofpain",
		ID:            39,
		LocalizedName: "Queen of Pain",
	},
	Hero{
		Name:          "venomancer",
		ID:            40,
		LocalizedName: "Venomancer",
	},
	Hero{
		Name:          "faceless_void",
		ID:            41,
		LocalizedName: "Faceless Void",
	},
	Hero{
		Name:          "skeleton_king",
		ID:            42,
		LocalizedName: "Skeleton King",
	},
	Hero{
		Name:          "death_prophet",
		ID:            43,
		LocalizedName: "Death Prophet",
	},
	Hero{
		Name:          "phantom_assassin",
		ID:            44,
		LocalizedName: "Phantom Assassin",
	},
	Hero{
		Name:          "pugna",
		ID:            45,
		LocalizedName: "Pugna",
	},
	Hero{
		Name:          "templar_assassin",
		ID:            46,
		LocalizedName: "Templar Assassin",
	},
	Hero{
		Name:          "viper",
		ID:            47,
		LocalizedName: "Viper",
	},
	Hero{
		Name:          "luna",
		ID:            48,
		LocalizedName: "Luna",
	},
	Hero{
		Name:          "dragon_knight",
		ID:            49,
		LocalizedName: "Dragon Knight",
	},
	Hero{
		Name:          "dazzle",
		ID:            50,
		LocalizedName: "Dazzle",
	},
	Hero{
		Name:          "rattletrap",
		ID:            51,
		LocalizedName: "Clockwerk",
	},
	Hero{
		Name:          "leshrac",
		ID:            52,
		LocalizedName: "Leshrac",
	},
	Hero{
		Name:          "furion",
		ID:            53,
		LocalizedName: "Nature's Prophet",
	},
	Hero{
		Name:          "life_stealer",
		ID:            54,
		LocalizedName: "Lifestealer",
	},
	Hero{
		Name:          "dark_seer",
		ID:            55,
		LocalizedName: "Dark Seer",
	},
	Hero{
		Name:          "clinkz",
		ID:            56,
		LocalizedName: "Clinkz",
	},
	Hero{
		Name:          "omniknight",
		ID:            57,
		LocalizedName: "Omniknight",
	},
	Hero{
		Name:          "enchantress",
		ID:            58,
		LocalizedName: "Enchantress",
	},
	Hero{
		Name:          "huskar",
		ID:            59,
		LocalizedName: "Huskar",
	},
	Hero{
		Name:          "night_stalker",
		ID:            60,
		LocalizedName: "Night Stalker",
	},
	Hero{
		Name:          "broodmother",
		ID:            61,
		LocalizedName: "Broodmother",
	},
	Hero{
		Name:          "bounty_hunter",
		ID:            62,
		LocalizedName: "Bounty Hunter",
	},
	Hero{
		Name:          "weaver",
		ID:            63,
		LocalizedName: "Weaver",
	},
	Hero{
		Name:          "jakiro",
		ID:            64,
		LocalizedName: "Jakiro",
	},
	Hero{
		Name:          "batrider",
		ID:            65,
		LocalizedName: "Batrider",
	},
	Hero{
		Name:          "chen",
		ID:            66,
		LocalizedName: "Chen",
	},
	Hero{
		Name:          "spectre",
		ID:            67,
		LocalizedName: "Spectre",
	},
	Hero{
		Name:          "doom_bringer",
		ID:            69,
		LocalizedName: "Doom",
	},
	Hero{
		Name:          "ancient_apparition",
		ID:            68,
		LocalizedName: "Ancient Apparition",
	},
	Hero{
		Name:          "ursa",
		ID:            70,
		LocalizedName: "Ursa",
	},
	Hero{
		Name:          "spirit_breaker",
		ID:            71,
		LocalizedName: "Spirit Breaker",
	},
	Hero{
		Name:          "gyrocopter",
		ID:            72,
		LocalizedName: "Gyrocopter",
	},
	Hero{
		Name:          "alchemist",
		ID:            73,
		LocalizedName: "Alchemist",
	},
	Hero{
		Name:          "invoker",
		ID:            74,
		LocalizedName: "Invoker",
	},
	Hero{
		Name:          "silencer",
		ID:            75,
		LocalizedName: "Silencer",
	},
	Hero{
		Name:          "obsidian_destroyer",
		ID:            76,
		LocalizedName: "Outworld Devourer",
	},
	Hero{
		Name:          "lycan",
		ID:            77,
		LocalizedName: "Lycanthrope",
	},
	Hero{
		Name:          "brewmaster",
		ID:            78,
		LocalizedName: "Brewmaster",
	},
	Hero{
		Name:          "shadow_demon",
		ID:            79,
		LocalizedName: "Shadow Demon",
	},
	Hero{
		Name:          "lone_druid",
		ID:            80,
		LocalizedName: "Lone Druid",
	},
	Hero{
		Name:          "chaos_knight",
		ID:            81,
		LocalizedName: "Chaos Knight",
	},
	Hero{
		Name:          "meepo",
		ID:            82,
		LocalizedName: "Meepo",
	},
	Hero{
		Name:          "treant",
		ID:            83,
		LocalizedName: "Treant Protector",
	},
	Hero{
		Name:          "ogre_magi",
		ID:            84,
		LocalizedName: "Ogre Magi",
	},
	Hero{
		Name:          "undying",
		ID:            85,
		LocalizedName: "Undying",
	},
	Hero{
		Name:          "rubick",
		ID:            86,
		LocalizedName: "Rubick",
	},
	Hero{
		Name:          "disruptor",
		ID:            87,
		LocalizedName: "Disruptor",
	},
	Hero{
		Name:          "nyx_assassin",
		ID:            88,
		LocalizedName: "Nyx Assassin",
	},
	Hero{
		Name:          "naga_siren",
		ID:            89,
		LocalizedName: "Naga Siren",
	},
	Hero{
		Name:          "keeper_of_the_light",
		ID:            90,
		LocalizedName: "Keeper of the Light",
	},
	Hero{
		Name:          "wisp",
		ID:            91,
		LocalizedName: "Wisp",
	},
	Hero{
		Name:          "visage",
		ID:            92,
		LocalizedName: "Visage",
	},
	Hero{
		Name:          "slark",
		ID:            93,
		LocalizedName: "Slark",
	},
	Hero{
		Name:          "medusa",
		ID:            94,
		LocalizedName: "Medusa",
	},
	Hero{
		Name:          "troll_warlord",
		ID:            95,
		LocalizedName: "Troll Warlord",
	},
	Hero{
		Name:          "centaur",
		ID:            96,
		LocalizedName: "Centaur Warrunner",
	},
	Hero{
		Name:          "magnataur",
		ID:            97,
		LocalizedName: "Magnus",
	},
	Hero{
		Name:          "shredder",
		ID:            98,
		LocalizedName: "Timbersaw",
	},
	Hero{
		Name:          "bristleback",
		ID:            99,
		LocalizedName: "Bristleback",
	},
	Hero{
		Name:          "tusk",
		ID:            100,
		LocalizedName: "Tusk",
	},
	Hero{
		Name:          "skywrath_mage",
		ID:            101,
		LocalizedName: "Skywrath Mage",
	},
	Hero{
		Name:          "abaddon",
		ID:            102,
		LocalizedName: "Abaddon",
	},
	Hero{
		Name:          "elder_titan",
		ID:            103,
		LocalizedName: "Elder Titan",
	},
	Hero{
		Name:          "legion_commander",
		ID:            104,
		LocalizedName: "Legion Commander",
	},
	Hero{
		Name:          "ember_spirit",
		ID:            106,
		LocalizedName: "Ember Spirit",
	},
	Hero{
		Name:          "earth_spirit",
		ID:            107,
		LocalizedName: "Earth Spirit",
	},
	Hero{
		Name:          "abyssal_underlord",
		ID:            108,
		LocalizedName: "Abyssal Underlord",
	},
	Hero{
		Name:          "terrorblade",
		ID:            109,
		LocalizedName: "Terrorblade",
	},
	Hero{
		Name:          "phoenix",
		ID:            110,
		LocalizedName: "Phoenix",
	},
	Hero{
		Name:          "techies",
		ID:            105,
		LocalizedName: "Techies",
	},
	Hero{
		Name:          "oracle",
		ID:            111,
		LocalizedName: "Oracle",
	},
	Hero{
		Name:          "winter_wyvern",
		ID:            112,
		LocalizedName: "Winter Wyvern",
	},
}
