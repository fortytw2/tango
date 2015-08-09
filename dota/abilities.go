package dota

// An Ability is one of 500~ usable skills in Dota2
type Ability struct {
	ID   int    `json:"id"`
	Name string `json:"ability_name"`
}

// A LevelUp occurs when a player chooses to skill a particular skill...
type LevelUp struct {
	ID    int    `json:"id"`
	Level int    `json:"level"`
	Time  string `json:"time"`
	Name  string `json:"ability_name"`
}

// Abilities is a list of every ability in Dota 2
var Abilities = []Ability{
	Ability{
		Name: "centaur_khan_war_stomp",
		ID:   5295,
	},
	Ability{
		Name: "necronomicon_warrior_last_will",
		ID:   5200,
	},
	Ability{
		Name: "enigma_demonic_conversion",
		ID:   5147,
	},
	Ability{
		Name: "beastmaster_boar_poison",
		ID:   5171,
	},
	Ability{
		Name: "juggernaut_blade_fury",
		ID:   5028,
	},
	Ability{
		Name: "witch_doctor_voodoo_restoration",
	},
	Ability{
		Name: "bloodseeker_blood_bath",
		ID:   5016,
	},
	Ability{
		Name: "backdoor_protection_in_base",
		ID:   5351,
	},
	Ability{
		Name: "broodmother_spawn_spiderite",
		ID:   5283,
	},
	Ability{
		Name: "spirit_breaker_charge_of_darkness",
		ID:   5353,
	},
	Ability{
		Name: "chaos_knight_reality_rift",
		ID:   5427,
	},
	Ability{
		Name: "lina_laguna_blade",
		ID:   5043,
	},
	Ability{
		Name: "riki_backstab",
		ID:   5144,
	},
	Ability{
		Name: "forest_troll_high_priest_heal",
		ID:   5318,
	},
	Ability{
		Name: "warlock_upheaval",
		ID:   5164,
	},
	Ability{
		Name: "ogre_magi_frost_armor",
		ID:   5304,
	},
	Ability{
		Name: "stats",
		ID:   5002,
	},
	Ability{
		Name: "dazzle_poison_touch",
		ID:   5233,
	},
	Ability{
		Name: "templar_assassin_self_trap",
		ID:   5199,
	},
	Ability{
		Name: "silencer_glaives_of_wisdom",
		ID:   5378,
	},
	Ability{
		Name: "warlock_rain_of_chaos",
		ID:   5165,
	},
	Ability{
		Name: "sven_warcry",
		ID:   5096,
	},
	Ability{
		Name: "templar_assassin_meld",
		ID:   5195,
	},
	Ability{
		Name: "enchantress_impetus",
		ID:   5270,
	},
	Ability{
		Name: "earthshaker_enchant_totem",
		ID:   5024,
	},
	Ability{
		Name: "lone_druid_synergy",
		ID:   5414,
	},
	Ability{
		Name: "invoker_forge_spirit",
		ID:   5387,
	},
	Ability{
		Name: "chaos_knight_phantasm",
		ID:   5429,
	},
	Ability{
		Name: "satyr_hellcaller_unholy_aura",
		ID:   5317,
	},
	Ability{
		Name: "shadow_shaman_shackles",
		ID:   5080,
	},
	Ability{
		Name: "ogre_magi_bloodlust",
		ID:   5440,
	},
	Ability{
		Name: "shadow_demon_shadow_poison_release",
		ID:   5424,
	},
	Ability{
		Name: "blue_dragonspawn_sorcerer_evasion",
		ID:   5325,
	},
	Ability{
		Name: "antimage_mana_break",
		ID:   5003,
	},
	Ability{
		Name: "spirit_breaker_empowering_haste",
		ID:   5354,
	},
	Ability{
		Name: "satyr_hellcaller_shockwave",
		ID:   5316,
	},
	Ability{
		Name: "morphling_morph_agi",
		ID:   5055,
	},
	Ability{
		Name: "bounty_hunter_track",
		ID:   5288,
	},
	Ability{
		Name: "bloodseeker_bloodrage",
		ID:   5015,
	},
	Ability{
		Name: "obsidian_destroyer_astral_imprisonment",
		ID:   5392,
	},
	Ability{
		Name: "lich_dark_ritual",
		ID:   5136,
	},
	Ability{
		Name: "antimage_spell_shield",
		ID:   5005,
	},
	Ability{
		Name: "enchantress_natures_attendants",
		ID:   5269,
	},
	Ability{
		Name: "broodmother_incapacitating_bite",
		ID:   5281,
	},
	Ability{
		Name: "undying_flesh_golem",
		ID:   5447,
	},
	Ability{
		Name: "lone_druid_rabid",
		ID:   5413,
	},
	Ability{
		Name: "slardar_sprint",
		ID:   5114,
	},
	Ability{
		Name: "sandking_caustic_finale",
		ID:   5104,
	},
	Ability{
		Name: "nevermore_dark_lord",
		ID:   5063,
	},
	Ability{
		Name: "lycan_summon_wolves_invisibility",
		ID:   5500,
	},
	Ability{
		Name: "life_stealer_infest",
		ID:   5252,
	},
	Ability{
		Name: "lycan_feral_impulse",
		ID:   5397,
	},
	Ability{
		Name: "ursa_overpower",
		ID:   5358,
	},
	Ability{
		Name: "enraged_wildkin_toughness_aura",
		ID:   5313,
	},
	Ability{
		Name: "pudge_flesh_heap",
		ID:   5074,
	},
	Ability{
		Name: "witch_doctor_maledict",
		ID:   5140,
	},
	Ability{
		Name: "vengefulspirit_nether_swap",
		ID:   5125,
	},
	Ability{
		Name: "weaver_shukuchi",
		ID:   5290,
	},
	Ability{
		Name: "weaver_the_swarm",
		ID:   5289,
	},
	Ability{
		Name: "alchemist_acid_spray",
		ID:   5365,
	},
	Ability{
		Name: "phantom_assassin_stifling_dagger",
		ID:   5190,
	},
	Ability{
		Name: "sandking_epicenter",
		ID:   5105,
	},
	Ability{
		Name: "tornado_tempest",
		ID:   5310,
	},
	Ability{
		Name: "jakiro_liquid_fire",
		ID:   5299,
	},
	Ability{
		Name: "chen_hand_of_god",
		ID:   5331,
	},
	Ability{
		Name: "batrider_flaming_lasso",
		ID:   5323,
	},
	Ability{
		Name: "lone_druid_true_form",
		ID:   5416,
	},
	Ability{
		Name: "spirit_breaker_nether_strike",
		ID:   5356,
	},
	Ability{
		Name: "spectre_haunt",
		ID:   5337,
	},
	Ability{
		Name: "alchemist_chemical_rage",
		ID:   5369,
	},
	Ability{
		Name: "silencer_last_word",
		ID:   5379,
	},
	Ability{
		Name: "tinker_rearm",
		ID:   5153,
	},
	Ability{
		Name: "rattletrap_rocket_flare",
		ID:   5239,
	},
	Ability{
		Name: "omniknight_guardian_angel",
		ID:   5266,
	},
	Ability{
		Name: "leshrac_diabolic_edict",
		ID:   5242,
	},
	Ability{
		Name: "tiny_toss",
		ID:   5107,
	},
	Ability{
		Name: "razor_static_link",
		ID:   5083,
	},
	Ability{
		Name: "invoker_ice_wall",
		ID:   5389,
	},
	Ability{
		Name: "brewmaster_drunken_haze",
		ID:   5401,
	},
	Ability{
		Name: "batrider_flamebreak",
		ID:   5321,
	},
	Ability{
		Name: "dark_seer_ion_shell",
		ID:   5256,
	},
	Ability{
		Name: "queenofpain_scream_of_pain",
		ID:   5175,
	},
	Ability{
		Name: "phantom_assassin_blur",
		ID:   5192,
	},
	Ability{
		Name: "phantom_lancer_juxtapose",
		ID:   5067,
	},
	Ability{
		Name: "tidehunter_anchor_smash",
		ID:   5120,
	},
	Ability{
		Name: "axe_battle_hunger",
		ID:   5008,
	},
	Ability{
		Name: "pugna_life_drain",
		ID:   5189,
	},
	Ability{
		Name: "batrider_firefly",
		ID:   5322,
	},
	Ability{
		Name: "beastmaster_primal_roar",
		ID:   5177,
	},
	Ability{
		Name: "shadow_shaman_mass_serpent_ward",
		ID:   5081,
	},
	Ability{
		Name: "treant_leech_seed",
		ID:   5435,
	},
	Ability{
		Name: "chen_holy_persuasion",
		ID:   5330,
	},
	Ability{
		Name: "courier_burst",
		ID:   5210,
	},
	Ability{
		Name: "luna_lucent_beam",
		ID:   5222,
	},
	Ability{
		Name: "morphling_morph_replicate",
		ID:   5058,
	},
	Ability{
		Name: "venomancer_poison_sting",
		ID:   5179,
	},
	Ability{
		Name: "roshan_bash",
		ID:   5214,
	},
	Ability{
		Name: "tinker_march_of_the_machines",
		ID:   5152,
	},
	Ability{
		Name: "lone_druid_spirit_bear_demolish",
		ID:   5420,
	},
	Ability{
		Name: "mirana_leap",
		ID:   5050,
	},
	Ability{
		Name: "brewmaster_storm_wind_walk",
		ID:   5410,
	},
	Ability{
		Name: "skeleton_king_reincarnation",
		ID:   5089,
	},
	Ability{
		Name: "beastmaster_call_of_the_wild",
		ID:   5169,
	},
	Ability{
		Name: "riki_permanent_invisibility",
		ID:   5145,
	},
	Ability{
		Name: "courier_shield",
		ID:   5209,
	},
	Ability{
		Name: "clinkz_searing_arrows",
		ID:   5260,
	},
	Ability{
		Name: "puck_illusory_orb",
		ID:   5069,
	},
	Ability{
		Name: "forged_spirit_melting_strike",
		ID:   5388,
	},
	Ability{
		Name: "viper_viper_strike",
		ID:   5221,
	},
	Ability{
		Name: "furion_wrath_of_nature",
		ID:   5248,
	},
	Ability{
		Name: "silencer_curse_of_the_silent",
		ID:   5377,
	},
	Ability{
		Name: "lich_frost_nova",
		ID:   5134,
	},
	Ability{
		Name: "doom_bringer_doom",
		ID:   5342,
	},
	Ability{
		Name: "warlock_golem_flaming_fists",
		ID:   5166,
	},
	Ability{
		Name: "invoker_quas",
		ID:   5370,
	},
	Ability{
		Name: "invoker_invoke",
		ID:   5375,
	},
	Ability{
		Name: "dazzle_weave",
		ID:   5236,
	},
	Ability{
		Name: "courier_go_to_secretshop",
		ID:   5492,
	},
	Ability{
		Name: "meepo_geostrike",
		ID:   5432,
	},
	Ability{
		Name: "ogre_magi_multicast",
		ID:   5441,
	},
	Ability{
		Name: "crystal_maiden_frostbite",
		ID:   5127,
	},
	Ability{
		Name: "invoker_deafening_blast",
		ID:   5390,
	},
	Ability{
		Name: "crystal_maiden_freezing_field",
		ID:   5129,
	},
	Ability{
		Name: "lina_light_strike_array",
		ID:   5041,
	},
	Ability{
		Name: "dragon_knight_breathe_fire",
		ID:   5226,
	},
	Ability{
		Name: "earthshaker_fissure",
		ID:   5023,
	},
	Ability{
		Name: "witch_doctor_paralyzing_cask",
		ID:   5138,
	},
	Ability{
		Name: "centaur_khan_endurance_aura",
		ID:   5294,
	},
	Ability{
		Name: "gyrocopter_homing_missile",
		ID:   5362,
	},
	Ability{
		Name: "death_prophet_silence",
		ID:   5091,
	},
	Ability{
		Name: "beastmaster_wild_axes",
		ID:   5168,
	},
	Ability{
		Name: "doom_bringer_lvl_death",
		ID:   5341,
	},
	Ability{
		Name: "queenofpain_shadow_strike",
		ID:   5173,
	},
	Ability{
		Name: "tiny_avalanche",
		ID:   5106,
	},
	Ability{
		Name: "invoker_exort",
		ID:   5372,
	},
	Ability{
		Name: "lone_druid_true_form_battle_cry",
		ID:   5417,
	},
	Ability{
		Name: "tiny_craggy_exterior",
		ID:   5108,
	},
	Ability{
		Name: "antimage_mana_void",
		ID:   5006,
	},
	Ability{
		Name: "riki_smoke_screen",
		ID:   5142,
	},
	Ability{
		Name: "roshan_spell_block",
		ID:   5213,
	},
	Ability{
		Name: "juggernaut_healing_ward",
		ID:   5029,
	},
	Ability{
		Name: "ancient_apparition_ice_blast_release",
		ID:   5349,
	},
	Ability{
		Name: "obsidian_destroyer_arcane_orb",
		ID:   5391,
	},
	Ability{
		Name: "earthshaker_aftershock",
		ID:   5025,
	},
	Ability{
		Name: "satyr_trickster_purge",
		ID:   5314,
	},
	Ability{
		Name: "slardar_amplify_damage",
		ID:   5117,
	},
	Ability{
		Name: "dazzle_shallow_grave",
		ID:   5234,
	},
	Ability{
		Name: "tiny_grow",
		ID:   5109,
	},
	Ability{
		Name: "shadow_demon_demonic_purge",
		ID:   5425,
	},
	Ability{
		Name: "enigma_midnight_pulse",
		ID:   5148,
	},
	Ability{
		Name: "treant_living_armor",
		ID:   5436,
	},
	Ability{
		Name: "spectre_desolate",
		ID:   5335,
	},
	Ability{
		Name: "weaver_time_lapse",
		ID:   5292,
	},
	Ability{
		Name: "ancient_apparition_chilling_touch",
		ID:   5347,
	},
	Ability{
		Name: "brewmaster_storm_cyclone",
		ID:   5409,
	},
	Ability{
		Name: "vengefulspirit_command_aura",
		ID:   5123,
	},
	Ability{
		Name: "broodmother_spawn_spiderlings",
		ID:   5279,
	},
	Ability{
		Name: "bane_fiends_grip",
		ID:   5013,
	},
	Ability{
		Name: "gyrocopter_flak_cannon",
		ID:   5363,
	},
	Ability{
		Name: "sniper_take_aim",
		ID:   5156,
	},
	Ability{
		Name: "sven_great_cleave",
		ID:   5095,
	},
	Ability{
		Name: "death_prophet_exorcism",
		ID:   5093,
	},
	Ability{
		Name: "mirana_arrow",
		ID:   5048,
	},
	Ability{
		Name: "bounty_hunter_jinada",
		ID:   5286,
	},
	Ability{
		Name: "pudge_rot",
		ID:   5076,
	},
	Ability{
		Name: "chen_test_of_faith",
		ID:   5329,
	},
	Ability{
		Name: "jakiro_ice_path",
		ID:   5298,
	},
	Ability{
		Name: "enraged_wildkin_tornado",
		ID:   5312,
	},
	Ability{
		Name: "pugna_decrepify",
		ID:   5187,
	},
	Ability{
		Name: "dark_troll_warlord_raise_dead",
		ID:   5306,
	},
	Ability{
		Name: "drow_ranger_trueshot",
		ID:   5021,
	},
	Ability{
		Name: "vengefulspirit_wave_of_terror",
		ID:   5124,
	},
	Ability{
		Name: "bounty_hunter_wind_walk",
		ID:   5287,
	},
	Ability{
		Name: "nevermore_requiem",
		ID:   5064,
	},
	Ability{
		Name: "courier_transfer_items",
		ID:   5206,
	},
	Ability{
		Name: "viper_poison_attack",
		ID:   5218,
	},
	Ability{
		Name: "brewmaster_fire_permanent_immolation",
		ID:   5411,
	},
	Ability{
		Name: "alchemist_unstable_concoction",
		ID:   5366,
	},
	Ability{
		Name: "undying_soul_rip",
		ID:   5443,
	},
	Ability{
		Name: "roshan_devotion",
		ID:   5217,
	},
	Ability{
		Name: "zuus_thundergods_wrath",
		ID:   5113,
	},
	Ability{
		Name: "broodmother_spin_web",
		ID:   5280,
	},
	Ability{
		Name: "jakiro_dual_breath",
		ID:   5297,
	},
	Ability{
		Name: "luna_moon_glaive",
		ID:   5223,
	},
	Ability{
		Name: "roshan_slam",
		ID:   5215,
	},
	Ability{
		Name: "queenofpain_sonic_wave",
		ID:   5176,
	},
	Ability{
		Name: "sniper_assassinate",
		ID:   5157,
	},
	Ability{
		Name: "invoker_tornado",
		ID:   5382,
	},
	Ability{
		Name: "dark_troll_warlord_ensnare",
		ID:   5305,
	},
	Ability{
		Name: "courier_return_stash_items",
		ID:   5207,
	},
	Ability{
		Name: "batrider_sticky_napalm",
		ID:   5320,
	},
	Ability{
		Name: "ancient_apparition_ice_vortex",
		ID:   5346,
	},
	Ability{
		Name: "windrunner_shackleshot",
		ID:   5130,
	},
	Ability{
		Name: "lion_voodoo",
		ID:   5045,
	},
	Ability{
		Name: "undying_decay",
		ID:   5442,
	},
	Ability{
		Name: "tinker_heat_seeking_missile",
		ID:   5151,
	},
	Ability{
		Name: "lion_mana_drain",
		ID:   5046,
	},
	Ability{
		Name: "lone_druid_spirit_bear",
		ID:   5412,
	},
	Ability{
		Name: "shadow_shaman_voodoo",
		ID:   5079,
	},
	Ability{
		Name: "brewmaster_thunder_clap",
		ID:   5400,
	},
	Ability{
		Name: "meepo_divided_we_stand",
		ID:   5433,
	},
	Ability{
		Name: "big_thunder_lizard_frenzy",
		ID:   5333,
	},
	Ability{
		Name: "luna_eclipse",
		ID:   5225,
	},
	Ability{
		Name: "sandking_sand_storm",
		ID:   5103,
	},
	Ability{
		Name: "giant_wolf_critical_strike",
		ID:   5307,
	},
	Ability{
		Name: "brewmaster_storm_dispel_magic",
		ID:   5408,
	},
	Ability{
		Name: "big_thunder_lizard_slam",
		ID:   5332,
	},
	Ability{
		Name: "forest_troll_high_priest_mana_aura",
		ID:   5491,
	},
	Ability{
		Name: "invoker_wex",
		ID:   5371,
	},
	Ability{
		Name: "clinkz_wind_walk",
		ID:   5261,
	},
	Ability{
		Name: "lone_druid_true_form",
		ID:   5415,
	},
	Ability{
		Name: "undying_tombstone_zombie_deathstrike",
		ID:   5446,
	},
	Ability{
		Name: "queenofpain_blink",
		ID:   5174,
	},
	Ability{
		Name: "roshan_illusion_protection",
		ID:   5216,
	},
	Ability{
		Name: "necronomicon_archer_aoe",
		ID:   5204,
	},
	Ability{
		Name: "witch_doctor_death_ward",
		ID:   5141,
	},
	Ability{
		Name: "enchantress_untouchable",
		ID:   5267,
	},
	Ability{
		Name: "night_stalker_void",
		ID:   5275,
	},
	Ability{
		Name: "faceless_void_time_lock",
		ID:   5184,
	},
	Ability{
		Name: "faceless_void_backtrack",
		ID:   5183,
	},
	Ability{
		Name: "polar_furbolg_ursa_warrior_thunder_clap",
		ID:   5302,
	},
	Ability{
		Name: "slardar_bash",
		ID:   5116,
	},
	Ability{
		Name: "omniknight_repel",
		ID:   5264,
	},
	Ability{
		Name: "huskar_life_break",
		ID:   5274,
	},
	Ability{
		Name: "spirit_breaker_greater_bash",
		ID:   5355,
	},
	Ability{
		Name: "invoker_ghost_walk",
		ID:   5381,
	},
	Ability{
		Name: "tidehunter_gush",
		ID:   5118,
	},
	Ability{
		Name: "tidehunter_ravage",
		ID:   5121,
	},
	Ability{
		Name: "lycan_summon_wolves",
		ID:   5395,
	},
	Ability{
		Name: "drow_ranger_frost_arrows",
		ID:   5019,
	},
	Ability{
		Name: "rattletrap_power_cogs",
		ID:   5238,
	},
	Ability{
		Name: "chen_penitence",
		ID:   5328,
	},
	Ability{
		Name: "skeleton_king_critical_strike",
		ID:   5088,
	},
	Ability{
		Name: "dazzle_shadow_wave",
		ID:   5235,
	},
	Ability{
		Name: "sniper_headshot",
		ID:   5155,
	},
	Ability{
		Name: "life_stealer_consume",
		ID:   5253,
	},
	Ability{
		Name: "kunkka_x_marks_the_spot",
		ID:   5033,
	},
	Ability{
		Name: "furion_teleportation",
		ID:   5246,
	},
	Ability{
		Name: "brewmaster_drunken_brawler",
		ID:   5402,
	},
	Ability{
		Name: "morphling_waveform",
		ID:   5052,
	},
	Ability{
		Name: "morphling_replicate",
		ID:   5057,
	},
	Ability{
		Name: "storm_spirit_static_remnant",
		ID:   5098,
	},
	Ability{
		Name: "drow_ranger_marksmanship",
		ID:   5022,
	},
	Ability{
		Name: "night_stalker_darkness",
		ID:   5278,
	},
	Ability{
		Name: "windrunner_powershot",
		ID:   5131,
	},
	Ability{
		Name: "venomancer_plague_ward",
		ID:   5180,
	},
	Ability{
		Name: "invoker_sun_strike",
		ID:   5386,
	},
	Ability{
		Name: "riki_blink_strike",
		ID:   5143,
	},
	Ability{
		Name: "zuus_static_field",
		ID:   5112,
	},
	Ability{
		Name: "pudge_meat_hook",
		ID:   5075,
	},
	Ability{
		Name: "pugna_nether_ward",
		ID:   5188,
	},
	Ability{
		Name: "treant_overgrowth",
		ID:   5437,
	},
	Ability{
		Name: "zuus_lightning_bolt",
		ID:   5111,
	},
	Ability{
		Name: "luna_lunar_blessing",
		ID:   5224,
	},
	Ability{
		Name: "brewmaster_primal_split",
		ID:   5403,
	},
	Ability{
		Name: "invoker_chaos_meteor",
		ID:   5385,
	},
	Ability{
		Name: "ogre_magi_fireblast",
		ID:   5438,
	},
	Ability{
		Name: "tidehunter_kraken_shell",
		ID:   5119,
	},
	Ability{
		Name: "ursa_enrage",
		ID:   5360,
	},
	Ability{
		Name: "lich_chain_frost",
		ID:   5137,
	},
	Ability{
		Name: "razor_eye_of_the_storm",
		ID:   5085,
	},
	Ability{
		Name: "ursa_fury_swipes",
		ID:   5359,
	},
	Ability{
		Name: "juggernaut_omni_slash",
		ID:   5030,
	},
	Ability{
		Name: "obsidian_destroyer_essence_aura",
		ID:   5393,
	},
	Ability{
		Name: "warlock_shadow_word",
		ID:   5163,
	},
	Ability{
		Name: "axe_berserkers_call",
		ID:   5007,
	},
	Ability{
		Name: "dark_seer_wall_of_replica",
		ID:   5258,
	},
	Ability{
		Name: "razor_unstable_current",
		ID:   5084,
	},
	Ability{
		Name: "necrolyte_death_pulse",
		ID:   5158,
	},
	Ability{
		Name: "necrolyte_reapers_scythe",
		ID:   5161,
	},
	Ability{
		Name: "broodmother_insatiable_hunger",
		ID:   5282,
	},
	Ability{
		Name: "templar_assassin_psi_blades",
		ID:   5196,
	},
	Ability{
		Name: "pugna_nether_blast",
		ID:   5186,
	},
	Ability{
		Name: "leshrac_pulse_nova",
		ID:   5244,
	},
	Ability{
		Name: "doom_bringer_scorched_earth",
		ID:   5340,
	},
	Ability{
		Name: "warlock_golem_permanent_immolation",
		ID:   5167,
	},
	Ability{
		Name: "bloodseeker_thirst",
		ID:   5017,
	},
	Ability{
		Name: "mirana_starfall",
		ID:   5051,
	},
	Ability{
		Name: "viper_corrosive_skin",
		ID:   5220,
	},
	Ability{
		Name: "kunkka_ghostship",
		ID:   5035,
	},
	Ability{
		Name: "dark_seer_surge",
		ID:   5257,
	},
	Ability{
		Name: "lone_druid_spirit_bear_return",
		ID:   5418,
	},
	Ability{
		Name: "meepo_poof",
		ID:   5431,
	},
	Ability{
		Name: "black_dragon_splash_attack",
		ID:   5324,
	},
	Ability{
		Name: "rattletrap_battery_assault",
		ID:   5237,
	},
	Ability{
		Name: "drow_ranger_silence",
		ID:   5020,
	},
	Ability{
		Name: "nevermore_necromastery",
		ID:   5062,
	},
	Ability{
		Name: "jakiro_macropyre",
		ID:   5300,
	},
	Ability{
		Name: "lone_druid_spirit_bear_entangle",
		ID:   5419,
	},
	Ability{
		Name: "treant_natures_guise",
		ID:   5434,
	},
	Ability{
		Name: "beastmaster_greater_boar_poison",
		ID:   5352,
	},
	Ability{
		Name: "tinker_laser",
		ID:   5150,
	},
	Ability{
		Name: "clinkz_death_pact",
		ID:   5262,
	},
	Ability{
		Name: "morphling_morph",
		ID:   5054,
	},
	Ability{
		Name: "morphling_adaptive_strike",
		ID:   5053,
	},
	Ability{
		Name: "mirana_invis",
		ID:   5049,
	},
	Ability{
		Name: "enigma_malefice",
		ID:   5146,
	},
	Ability{
		Name: "skeleton_king_hellfire_blast",
		ID:   5086,
	},
	Ability{
		Name: "silencer_global_silence",
		ID:   5380,
	},
	Ability{
		Name: "phantom_assassin_phantom_strike",
		ID:   5191,
	},
	Ability{
		Name: "lycan_shapeshift",
		ID:   5398,
	},
	Ability{
		Name: "templar_assassin_refraction",
		ID:   5194,
	},
	Ability{
		Name: "warlock_fatal_bonds",
		ID:   5162,
	},
	Ability{
		Name: "death_prophet_carrion_swarm",
		ID:   5090,
	},
	Ability{
		Name: "lion_finger_of_death",
		ID:   5047,
	},
	Ability{
		Name: "ursa_earthshock",
		ID:   5357,
	},
	Ability{
		Name: "life_stealer_open_wounds",
		ID:   5251,
	},
	Ability{
		Name: "blue_dragonspawn_overseer_evasion",
		ID:   5326,
	},
	Ability{
		Name: "slardar_slithereen_crush",
		ID:   5115,
	},
	Ability{
		Name: "razor_plasma_field",
		ID:   5082,
	},
	Ability{
		Name: "storm_spirit_electric_vortex",
		ID:   5099,
	},
	Ability{
		Name: "spectre_dispersion",
		ID:   5336,
	},
	Ability{
		Name: "courier_return_to_base",
		ID:   5205,
	},
	Ability{
		Name: "sven_gods_strength",
		ID:   5097,
	},
	Ability{
		Name: "leshrac_lightning_storm",
		ID:   5243,
	},
	Ability{
		Name: "life_stealer_rage",
		ID:   5249,
	},
	Ability{
		Name: "spectre_reality",
		ID:   5338,
	},
	Ability{
		Name: "antimage_blink",
		ID:   5004,
	},
	Ability{
		Name: "sniper_shrapnel",
		ID:   5154,
	},
	Ability{
		Name: "rattletrap_hookshot",
		ID:   5240,
	},
	Ability{
		Name: "invoker_cold_snap",
		ID:   5376,
	},
	Ability{
		Name: "crystal_maiden_crystal_nova",
		ID:   5126,
	},
	Ability{
		Name: "bloodseeker_rupture",
		ID:   5018,
	},
	Ability{
		Name: "kunkka_tidebringer",
		ID:   5032,
	},
	Ability{
		Name: "invoker_alacrity",
		ID:   5384,
	},
	Ability{
		Name: "huskar_inner_vitality",
		ID:   5271,
	},
	Ability{
		Name: "axe_counter_helix",
		ID:   5009,
	},
	Ability{
		Name: "spectre_spectral_dagger",
		ID:   5334,
	},
	Ability{
		Name: "furion_force_of_nature",
		ID:   5247,
	},
	Ability{
		Name: "brewmaster_earth_hurl_boulder",
		ID:   5404,
	},
	Ability{
		Name: "alchemist_goblins_greed",
		ID:   5368,
	},
	Ability{
		Name: "windrunner_windrun",
		ID:   5132,
	},
	Ability{
		Name: "kunkka_return",
		ID:   5034,
	},
	Ability{
		Name: "backdoor_protection",
		ID:   5350,
	},
	Ability{
		Name: "dragon_knight_dragon_blood",
		ID:   5228,
	},
	Ability{
		Name: "vengefulspirit_magic_missile",
		ID:   5122,
	},
	Ability{
		Name: "phantom_lancer_spirit_lance",
		ID:   5065,
	},
	Ability{
		Name: "nevermore_shadowraze1",
		ID:   5059,
	},
	Ability{
		Name: "courier_take_stash_items",
		ID:   5208,
	},
	Ability{
		Name: "nevermore_shadowraze2",
		ID:   5060,
	},
	Ability{
		Name: "nevermore_shadowraze3",
		ID:   5061,
	},
	Ability{
		Name: "weaver_geminate_attack",
		ID:   5291,
	},
	Ability{
		Name: "alpha_wolf_command_aura",
		ID:   5309,
	},
	Ability{
		Name: "ancient_apparition_ice_blast",
		ID:   5348,
	},
	Ability{
		Name: "ogre_magi_ignite",
		ID:   5439,
	},
	Ability{
		Name: "harpy_storm_chain_lightning",
		ID:   5319,
	},
	Ability{
		Name: "omniknight_degen_aura",
		ID:   5265,
	},
	Ability{
		Name: "enchantress_enchant",
		ID:   5268,
	},
	Ability{
		Name: "leshrac_split_earth",
		ID:   5241,
	},
	Ability{
		Name: "death_prophet_witchcraft",
		ID:   5092,
	},
	Ability{
		Name: "viper_nethertoxin",
		ID:   5219,
	},
	Ability{
		Name: "necrolyte_sadist",
		ID:   5160,
	},
	Ability{
		Name: "phantom_assassin_coup_de_grace",
		ID:   5193,
	},
	Ability{
		Name: "ancient_apparition_cold_feet",
		ID:   5345,
	},
	Ability{
		Name: "puck_dream_coil",
		ID:   5073,
	},
	Ability{
		Name: "venomancer_poison_nova",
		ID:   5181,
	},
	Ability{
		Name: "chaos_knight_chaos_bolt",
		ID:   5426,
	},
	Ability{
		Name: "doom_bringer_devour",
		ID:   5339,
	},
	Ability{
		Name: "bounty_hunter_shuriken_toss",
		ID:   5285,
	},
	Ability{
		Name: "blue_dragonspawn_overseer_devotion_aura",
		ID:   5327,
	},
	Ability{
		Name: "earthshaker_echo_slam",
		ID:   5026,
	},
	Ability{
		Name: "necronomicon_warrior_sight",
		ID:   5201,
	},
	Ability{
		Name: "huskar_berserkers_blood",
		ID:   5273,
	},
	Ability{
		Name: "templar_assassin_psionic_trap",
		ID:   5197,
	},
	Ability{
		Name: "night_stalker_crippling_fear",
		ID:   5276,
	},
	Ability{
		Name: "storm_spirit_overload",
		ID:   5100,
	},
	Ability{
		Name: "alchemist_unstable_concoction_throw",
		ID:   5367,
	},
	Ability{
		Name: "morphling_morph_str",
		ID:   5056,
	},
	Ability{
		Name: "doom_bringer_empty",
		ID:   5343,
	},
	Ability{
		Name: "doom_bringer_empty",
		ID:   5344,
	},
	Ability{
		Name: "lycan_howl",
		ID:   5396,
	},
	Ability{
		Name: "beastmaster_hawk_invisibility",
		ID:   5170,
	},
	Ability{
		Name: "storm_spirit_ball_lightning",
		ID:   5101,
	},
	Ability{
		Name: "omniknight_purification",
		ID:   5263,
	},
	Ability{
		Name: "shadow_demon_shadow_poison",
		ID:   5423,
	},
	Ability{
		Name: "dragon_knight_elder_dragon_form",
		ID:   5229,
	},
	Ability{
		Name: "juggernaut_blade_dance",
		ID:   5027,
	},
	Ability{
		Name: "shadow_demon_disruption",
		ID:   5421,
	},
	Ability{
		Name: "venomancer_venomous_gale",
		ID:   5178,
	},
	Ability{
		Name: "gnoll_assassin_envenomed_weapon",
		ID:   5296,
	},
	Ability{
		Name: "alpha_wolf_critical_strike",
		ID:   5308,
	},
	Ability{
		Name: "shadow_shaman_ether_shock",
		ID:   5078,
	},
	Ability{
		Name: "night_stalker_hunter_in_the_night",
		ID:   5277,
	},
	Ability{
		Name: "dragon_knight_dragon_tail",
		ID:   5227,
	},
	Ability{
		Name: "gyrocopter_call_down",
		ID:   5364,
	},
	Ability{
		Name: "beastmaster_inner_beast",
		ID:   5172,
	},
	Ability{
		Name: "pudge_dismember",
		ID:   5077,
	},
	Ability{
		Name: "windrunner_focusfire",
		ID:   5133,
	},
	Ability{
		Name: "ghost_frost_attack",
		ID:   5301,
	},
	Ability{
		Name: "bane_brain_sap",
		ID:   5011,
	},
	Ability{
		Name: "necrolyte_heartstopper_aura",
		ID:   5159,
	},
	Ability{
		Name: "chaos_knight_chaos_strike",
		ID:   5428,
	},
	Ability{
		Name: "skeleton_king_vampiric_aura",
		ID:   5087,
	},
	Ability{
		Name: "clinkz_strafe",
		ID:   5259,
	},
	Ability{
		Name: "life_stealer_feast",
		ID:   5250,
	},
	Ability{
		Name: "faceless_void_chronosphere",
		ID:   5185,
	},
	Ability{
		Name: "lion_impale",
		ID:   5044,
	},
	Ability{
		Name: "necronomicon_archer_mana_burn",
		ID:   5203,
	},
	Ability{
		Name: "axe_culling_blade",
		ID:   5010,
	},
	Ability{
		Name: "huskar_burning_spear",
		ID:   5272,
	},
	Ability{
		Name: "brewmaster_earth_pulverize",
		ID:   5406,
	},
	Ability{
		Name: "kobold_taskmaster_speed_aura",
		ID:   5293,
	},
	Ability{
		Name: "brewmaster_earth_spell_immunity",
		ID:   5405,
	},
	Ability{
		Name: "necronomicon_warrior_mana_burn",
		ID:   5202,
	},
	Ability{
		Name: "furion_sprout",
		ID:   5245,
	},
	Ability{
		Name: "invoker_empty2",
		ID:   5374,
	},
	Ability{
		Name: "puck_ethereal_jaunt",
		ID:   5070,
	},
	Ability{
		Name: "invoker_empty1",
		ID:   5373,
	},
	Ability{
		Name: "kunkka_torrent",
		ID:   5031,
	},
	Ability{
		Name: "phantom_lancer_doppelwalk",
		ID:   5066,
	},
	Ability{
		Name: "undying_tombstone",
		ID:   5444,
	},
	Ability{
		Name: "bane_nightmare",
		ID:   5014,
	},
	Ability{
		Name: "zuus_arc_lightning",
		ID:   5110,
	},
	Ability{
		Name: "undying_tombstone_zombie_aura",
		ID:   5445,
	},
	Ability{
		Name: "crystal_maiden_brilliance_aura",
		ID:   5128,
	},
	Ability{
		Name: "bane_enfeeble",
		ID:   5012,
	},
	Ability{
		Name: "satyr_soulstealer_mana_burn",
		ID:   5315,
	},
	Ability{
		Name: "gyrocopter_rocket_barrage",
		ID:   5361,
	},
	Ability{
		Name: "dragon_knight_frost_breath",
		ID:   5232,
	},
	Ability{
		Name: "invoker_emp",
		ID:   5383,
	},
	Ability{
		Name: "shadow_demon_soul_catcher",
		ID:   5422,
	},
	Ability{
		Name: "lina_fiery_soul",
		ID:   5042,
	},
	Ability{
		Name: "lina_dragon_slave",
		ID:   5040,
	},
	Ability{
		Name: "sven_storm_bolt",
		ID:   5094,
	},
	Ability{
		Name: "broodmother_poison_sting",
		ID:   5284,
	},
	Ability{
		Name: "meepo_earthbind",
		ID:   5430,
	},
	Ability{
		Name: "phantom_lancer_phantom_edge",
		ID:   5068,
	},
	Ability{
		Name: "puck_phase_shift",
		ID:   5072,
	},
	Ability{
		Name: "faceless_void_time_walk",
		ID:   5182,
	},
	Ability{
		Name: "puck_waning_rift",
		ID:   5071,
	},
	Ability{
		Name: "lich_frost_armor",
		ID:   5135,
	},
	Ability{
		Name: "sandking_burrowstrike",
		ID:   5102,
	},
	Ability{
		Name: "dark_seer_vacuum",
		ID:   5255,
	},
	Ability{
		Name: "obsidian_destroyer_sanity_eclipse",
		ID:   5394,
	},
	Ability{
		Name: "enigma_black_hole",
		ID:   5149,
	},
	Ability{
		Name: "lycan_summon_wolves_critical_strike",
		ID:   5399,
	},
	Ability{
		Name: "default_attack",
		ID:   5001,
	},
	Ability{
		Name: "templar_assassin_trap",
		ID:   5198,
	},
	Ability{
		Name: "neutral_spell_immunity",
		ID:   5303,
	},
	Ability{
		Name: "nyx_assassin_impale",
		ID:   5462,
	},
	Ability{
		Name: "nyx_assassin_mana_burn",
		ID:   5463,
	},
	Ability{
		Name: "nyx_assassin_spiked_carapace",
		ID:   5464,
	},
	Ability{
		Name: "nyx_assassin_vendetta",
		ID:   5465,
	},
	Ability{
		Name: "slark_dark_pact",
		ID:   5494,
	},
	Ability{
		Name: "slark_pounce",
		ID:   5495,
	},
	Ability{
		Name: "slark_essence_shift",
		ID:   5496,
	},
	Ability{
		Name: "slark_shadow_dance",
		ID:   5497,
	},
	Ability{
		Name: "medusa_split_shot",
		ID:   5504,
	},
	Ability{
		Name: "medusa_mystic_snake",
		ID:   5505,
	},
	Ability{
		Name: "medusa_mana_shield",
		ID:   5506,
	},
	Ability{
		Name: "medusa_stone_gaze",
		ID:   5507,
	},
	Ability{
		Name: "troll_warlord_berserkers_rage",
		ID:   5508,
	},
	Ability{
		Name: "troll_warlord_whirling_axes_ranged",
		ID:   5509,
	},
	Ability{
		Name: "troll_warlord_whirling_axes_melee",
		ID:   5510,
	},
	Ability{
		Name: "troll_warlord_fervor",
		ID:   5511,
	},
	Ability{
		Name: "troll_warlord_battle_trance",
		ID:   5512,
	},
	Ability{
		Name: "centaur_hoof_stomp",
		ID:   5514,
	},
	Ability{
		Name: "centaur_double_edge",
		ID:   5515,
	},
	Ability{
		Name: "centaur_return",
		ID:   5516,
	},
	Ability{
		Name: "centaur_stampede",
		ID:   5517,
	},
	Ability{
		Name: "magnataur_shockwave",
		ID:   5518,
	},
	Ability{
		Name: "magnataur_empower",
		ID:   5519,
	},
	Ability{
		Name: "magnataur_skewer",
		ID:   5520,
	},
	Ability{
		Name: "magnataur_reverse_polarity",
		ID:   5521,
	},
	Ability{
		Name: "shredder_whirling_death",
		ID:   5524,
	},
	Ability{
		Name: "shredder_timber_chain",
		ID:   5525,
	},
	Ability{
		Name: "shredder_reactive_armor",
		ID:   5526,
	},
	Ability{
		Name: "shredder_chakram",
		ID:   5527,
	},
	Ability{
		Name: "shredder_return_chakram",
		ID:   5528,
	},
	Ability{
		Name: "shredder_chakram_2",
		ID:   5645,
	},
	Ability{
		Name: "shredder_return_chakram_2",
		ID:   5646,
	},
	Ability{
		Name: "keeper_of_the_light_illuminate",
		ID:   5471,
	},
	Ability{
		Name: "keeper_of_the_light_mana_leak",
		ID:   5472,
	},
	Ability{
		Name: "keeper_of_the_light_chakra_magic",
		ID:   5473,
	},
	Ability{
		Name: "keeper_of_the_light_spirit_form",
		ID:   5474,
	},
	Ability{
		Name: "keeper_of_the_light_recall",
		ID:   5475,
	},
	Ability{
		Name: "keeper_of_the_light_blinding_light",
		ID:   5476,
	},
	Ability{
		Name: "keeper_of_the_light_illuminate_end",
		ID:   5477,
	},
	Ability{
		Name: "keeper_of_the_light_spirit_form_illuminate",
		ID:   5479,
	},
	Ability{
		Name: "keeper_of_the_light_spirit_form_illuminate_end",
		ID:   5503,
	},
	Ability{
		Name: "rubick_telekinesis",
		ID:   5448,
	},
	Ability{
		Name: "rubick_fade_bolt",
		ID:   5450,
	},
	Ability{
		Name: "rubick_null_field",
		ID:   5451,
	},
	Ability{
		Name: "rubick_spell_steal",
		ID:   5452,
	},
	Ability{
		Name: "disruptor_thunder_strike",
		ID:   5458,
	},
	Ability{
		Name: "disruptor_glimpse",
		ID:   5459,
	},
	Ability{
		Name: "disruptor_kinetic_field",
		ID:   5460,
	},
	Ability{
		Name: "disruptor_static_storm",
		ID:   5461,
	},
	Ability{
		Name: "naga_siren_mirror_image",
		ID:   5467,
	},
	Ability{
		Name: "naga_siren_ensnare",
		ID:   5468,
	},
	Ability{
		Name: "naga_siren_rip_tide",
		ID:   5469,
	},
	Ability{
		Name: "naga_siren_song_of_the_siren",
		ID:   5470,
	},
	Ability{
		Name: "visage_grave_chill",
		ID:   5480,
	},
	Ability{
		Name: "visage_soul_assumption",
		ID:   5481,
	},
	Ability{
		Name: "visage_gravekeepers_cloak",
		ID:   5482,
	},
	Ability{
		Name: "visage_summon_familiars",
		ID:   5483,
	},
	Ability{
		Name: "wisp_tether",
		ID:   5485,
	},
	Ability{
		Name: "wisp_spirits",
		ID:   5486,
	},
	Ability{
		Name: "wisp_overcharge",
		ID:   5487,
	},
	Ability{
		Name: "wisp_relocate",
		ID:   5488,
	},
	Ability{
		Name: "wisp_tether_break",
		ID:   5489,
	},
	Ability{
		Name: "wisp_spirits_in",
		ID:   5490,
	},
	Ability{
		Name: "wisp_spirits_out",
		ID:   5493,
	},
	Ability{
		Name: "bristleback_viscous_nasal_goo",
		ID:   5548,
	},
	Ability{
		Name: "bristleback_quill_spray",
		ID:   5549,
	},
	Ability{
		Name: "bristleback_bristleback",
		ID:   5550,
	},
	Ability{
		Name: "bristleback_warpath",
		ID:   5551,
	},
	Ability{
		Name: "tusk_ice_shards",
		ID:   5565,
	},
	Ability{
		Name: "tusk_snowball",
		ID:   5566,
	},
	Ability{
		Name: "tusk_frozen_sigil",
		ID:   5567,
	},
	Ability{
		Name: "tusk_walrus_punch",
		ID:   5568,
	},
	Ability{
		Name: "skywrath_mage_arcane_bolt",
		ID:   5581,
	},
	Ability{
		Name: "skywrath_mage_concussive_shot",
		ID:   5582,
	},
	Ability{
		Name: "skywrath_mage_ancient_seal",
		ID:   5583,
	},
	Ability{
		Name: "skywrath_mage_mystic_flare",
		ID:   5584,
	},
	Ability{
		Name: "abaddon_death_coil",
		ID:   5585,
	},
	Ability{
		Name: "abaddon_aphotic_shield",
		ID:   5586,
	},
	Ability{
		Name: "abaddon_frostmourne",
		ID:   5587,
	},
	Ability{
		Name: "abaddon_borrowed_time",
		ID:   5588,
	},
	Ability{
		Name: "elder_titan_echo_stomp",
		ID:   5589,
	},
	Ability{
		Name: "elder_titan_ancestral_spirit",
		ID:   5591,
	},
	Ability{
		Name: "elder_titan_natural_order",
		ID:   5593,
	},
	Ability{
		Name: "elder_titan_earth_splitter",
		ID:   5594,
	},
	Ability{
		Name: "legion_commander_overwhelming_odds",
		ID:   5595,
	},
	Ability{
		Name: "legion_commander_press_the_attack",
		ID:   5596,
	},
	Ability{
		Name: "legion_commander_moment_of_courage",
		ID:   5597,
	},
	Ability{
		Name: "legion_commander_duel",
		ID:   5598,
	},
	Ability{
		Name: "techies_land_mines",
		ID:   5599,
	},
	Ability{
		Name: "techies_stasis_trap",
		ID:   5600,
	},
	Ability{
		Name: "techies_suicide_squad_attack",
		ID:   5601,
	},
	Ability{
		Name: "techies_remote_mines",
		ID:   5602,
	},
	Ability{
		Name: "techies_focused_detonate",
		ID:   5635,
	},
	Ability{
		Name: "techies_remote_mines_self_detonate",
		ID:   5636,
	},
	Ability{
		Name: "techies_minefield_sign",
		ID:   5644,
	},
	Ability{
		Name: "ember_spirit_searing_chains",
		ID:   5603,
	},
	Ability{
		Name: "ember_spirit_sleight_of_fist",
		ID:   5604,
	},
	Ability{
		Name: "ember_spirit_flame_guard",
		ID:   5605,
	},
	Ability{
		Name: "ember_spirit_fire_remnant",
		ID:   5606,
	},
	Ability{
		Name: "ember_spirit_activate_fire_remnant",
		ID:   5607,
	},
	Ability{
		Name: "earth_spirit_boulder_smash",
		ID:   5608,
	},
	Ability{
		Name: "earth_spirit_rolling_boulder",
		ID:   5609,
	},
	Ability{
		Name: "earth_spirit_geomagnetic_grip",
		ID:   5610,
	},
	Ability{
		Name: "earth_spirit_stone_caller",
		ID:   5611,
	},
	Ability{
		Name: "earth_spirit_magnetize",
		ID:   5612,
	},
	Ability{
		Name: "abyssal_underlord_firestorm",
		ID:   5613,
	},
	Ability{
		Name: "abyssal_underlord_pit_of_malice",
		ID:   5614,
	},
	Ability{
		Name: "abyssal_underlord_atrophy_aura",
		ID:   5615,
	},
	Ability{
		Name: "abyssal_underlord_dark_rift",
		ID:   5616,
	},
	Ability{
		Name: "terrorblade_reflection",
		ID:   5619,
	},
	Ability{
		Name: "terrorblade_conjure_image",
		ID:   5620,
	},
	Ability{
		Name: "terrorblade_metamorphosis",
		ID:   5621,
	},
	Ability{
		Name: "terrorblade_sunder",
		ID:   5622,
	},
	Ability{
		Name: "phoenix_icarus_dive",
		ID:   5623,
	},
	Ability{
		Name: "phoenix_fire_spirits",
		ID:   5625,
	},
	Ability{
		Name: "phoenix_sun_ray",
		ID:   5626,
	},
	Ability{
		Name: "phoenix_supernova",
		ID:   5630,
	},
	Ability{
		Name: "drow_ranger_wave_of_silence",
		ID:   5632,
	},
	Ability{
		Name: "broodmother_spin_web_destroy",
		ID:   5643,
	},
	Ability{
		Name: "earth_spirit_petrify",
		ID:   5648,
	},
	Ability{
		Name: "treant_eyes_in_the_forest",
		ID:   5649,
	},
	Ability{
		Name: "beastmaster_call_of_the_wild_boar",
		ID:   5580,
	},
	Ability{
		Name: "life_stealer_control",
		ID:   5655,
	},
	Ability{
		Name: "oracle_fortunes_end",
		ID:   5637,
	},
	Ability{
		Name: "oracle_fates_edict",
		ID:   5638,
	},
	Ability{
		Name: "oracle_purifying_flames",
		ID:   5639,
	},
	Ability{
		Name: "oracle_false_promise",
		ID:   5640,
	},
	Ability{
		Name: "winter_wyvern_arctic_burn",
		ID:   5651,
	},
	Ability{
		Name: "winter_wyvern_splinter_blast",
		ID:   5652,
	},
	Ability{
		Name: "winter_wyvern_cold_embrace",
		ID:   5653,
	},
	Ability{
		Name: "winter_wyvern_winters_curse",
		ID:   5654,
	},
}
