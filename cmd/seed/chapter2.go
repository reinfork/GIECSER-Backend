// Chapter 2 content frozen from GIECSER_FINAL Lesson 2 (pp.24-43):
// goal setting → independent speaking. Self-reflection questions and the
// repractice loop live in the frontend — no rows by design.
package main

import "asri-backend/internal/model"

func ch2Final() seedChapter {
	living := `Deep within Riau's lush tropical forests and home gardens lies a living pharmacy. For generations, the Malay people have looked to the earth — not modern medicine cabinets — to heal, recharge, and stay strong. From bitter roots pulled straight from the soil to fragrant, sun-dried leaves, almost every plant tells a story of survival and care. Herbs like Pegagan and Akar Tunjuk Langit are gently brewed into earthy teas or crushed into healing poultices to treat everyday sickness, soothe fever, and restore vitality. It is more than just folk medicine; it is an enduring way of life built on patience, deep observation, and a profound gratitude for the natural world.`
	plants := `Pandan Nipah Jeruju Api-api Bakau Biduri Pegagan Mengkudu Limau Purut Akar pinang Beluntas Sirih Akar tunjuk langit Kupang-kupang`
	coreWords := `Root Sap Wound Tonic Heal Reduce Relieve Fragrant Antiseptic Bumpy`
	langFunc := `Identifying and naming: This plant is called pegagan. It is locally known in Riau as kupang-kupang. Describing appearance: It has smooth leaves. The taste is slightly bitter. Explaining uses: It is commonly used to treat wounds. People consume it to relieve back pain. It helps prevent you from bacteria. Recommendations: What should I drink to reduce my fever? You should try akar tunjuk langit because it has many benefits. It is good to chew Sirih leaves when you have a sore throat.`
	grammar := `Nominal sentence: Akar Tunjuk Langit is a rare terrestrial fern. Verbal sentence: This plant grows in damp, shaded secondary forests.`
	langExpr := `Heal: The crushed leaves help heal open cuts much faster. Relieve: Drinking this warm tonic helps relieve headache and body fatigue. Reduce: Consuming noni juice regularly can reduce high blood pressure. Soothe: The aroma of fresh pandan helps soothe nervous tension. Boil: First, boil five fresh leaves in two cups of water until it simmers. Fragrant: You can easily identify this plant by its sweet, fragrant leaves. Antiseptic: Betel leaf possesses strong antiseptic qualities to treat wounds. Bumpy: Kaffir lime is recognizable by its rough, bumpy green skin. Bitter: Although the herbal drink tastes bitter, it brings remarkable health benefits. Brackish: Nipah palms prefer the brackish waters of coastal river mouths.`
	putri := `Putri: Assalamu'alaikum, Pak Cik. Do you have a moment? I want to ask you about the traditional herbal medicine of our village for my school assignment. Pak Cik Rahim: Wa'alaikumussalam, Putri. Come, take a seat on the veranda. What kind of plant are you curious about? Putri: My grandfather mentioned a plant called Akar Tunjuk Langit. What kind of plant is it, Pak Cik? Why does it have such a curious name? Pak Cik Rahim: Alhamdulillah, it is good that youngsters like you still want to learn our heritage. It is a rare wild fern. We call it Tunjuk Langit — meaning pointing to the sky — because its central spike stands straight up, pointing right toward the heavens. Putri: Where does it usually grow? Can I find it easily behind my house? Pak Cik Rahim: Not in an open backyard, Putri. It likes quiet, damp places. It thrives deep in the shady secondary forests, near our freshwater swamp edges, and along moist riverbanks under the canopy of large trees. Putri: Which part do our people use as medicine? Pak Cik Rahim: The most valuable part is hidden underground: the root. We dig it up carefully, wash away the dark mud, slice it, and boil it in clean water. Putri: What are the health benefits of drinking the boiled root, Pak Cik? Pak Cik Rahim: It serves as a powerful traditional tonic. When fishermen or rubber tappers feel exhausted after working all day, drinking it restores their vitality and stamina. It is also famous across Riau for relieving stiff joints and bad lower back pain. Putri: Subhanallah, our Riau forests really are full of natural blessings. Thank you so much for explaining this, Pak Cik! Pak Cik Rahim: You are most welcome, Putri. Study well, and never forget the wisdom of our ancestors.`
	descriptive := `Akar Tunjuk Langit is a traditional plant that grows in some areas of Riau. Its scientific name is Helminthostachys zeylanica. The plant usually grows in moist and shady places, such as near forests and rivers. The Tunjuk Langit plant is a small green plant. It has long green leaves that grow close to the ground. The plant also has a special upright part that looks like a small stick pointing to the sky. This special shape is the reason why people call it Akar Tunjuk Langit, which means pointing to the sky. For generations, some Malay communities have known Tunjuk Langit as a traditional medicinal plant. The roots of the plant have traditionally been used as an ingredient in herbal preparations.`
	matchText := `The plant is called Tunjuk Langit because its upright spike stands straight toward the sky. Unlike open sunny fields, this plant naturally grows where it thrives in damp forest floors and freshwater swamp edges. The most useful part taken from this fern is its underground root system. Local people boil the root to make an herbal tonic that boosts stamina and restores physical energy. The natural properties of this plant also help relieve backache, stiff joints, and body soreness.`
	akarMono := `Good morning, everyone. Today, I want to describe a famous medicinal plant from Riau called Akar Tunjuk Langit. This plant gets its unique name because its upright stem points straight toward the sky. It is a primitive terrestrial fern with dark green leaves and a thick, fibrous root system. It thrives naturally in moist, shaded environments like tropical forest floors and freshwater swamp edges. The most valuable part of this plant is its root. In traditional Riau Malay medicine, the root is boiled to create a restorative tonic that boosts body stamina and energy. It is also well-known for helping relieve lower backache and joint pain.`
	mayaGap := `Maya: Nenek, what are you picking in the backyard? Those green plants smell very fresh! Nek Minah: I am picking some fresh sirih (1) _______ and pandan. They are useful traditional medicinal plants. Maya: Oh, I know pandan! It has a very sweet and (2) _______ aroma. What do you use it for? Nek Minah: We can (3) _______ the fresh leaves in hot water to make tea. Drinking it helps to calm your mind and (4) _______ high blood pressure. Maya: That is amazing! What about the sirih leaves? Nek Minah: Sirih leaf is famous for its strong (5) _______ properties that kill germs and bacteria. Yesterday, your uncle cut his finger, so I crushed the leaf and placed it directly over the open (6) _______. Maya: Does it help the cut to (7) _______ faster? Nek Minah: Yes, it stops the bleeding quickly and helps (8) _______ the pain and swelling. Maya: Traditional plants are really wonderful natural medicines, Nenek!`
	guideVocab := `Boil: You should boil the fresh leaves in water for ten minutes to extract their medicinal properties. Boost: Drinking fresh pegagan juice helps to boost your memory and brain focus. Grow: These wild medicinal herbs grow naturally along shaded riverbanks and paddy fields. Crush: Carefully crush the clean leaves with a mortar and apply the paste directly to the wound. Moist: Pegagan prefers moist soil with plenty of organic matter rather than dry, dusty earth. Healthy: Adding herbal plants to your daily diet is a simple step toward a healthy lifestyle. Raw: Many elders in Riau enjoy eating raw pegagan leaves as a fresh side vegetable with their rice. Fresh: Make sure to pick fresh green leaves in the morning before the sun gets too hot.`
	sirihDialog := `Student A: What is the plant in your garden? Student B: This is Sirih Leaves. It is a wild fern from Riau. Student A: Where does it usually grow? Student B: It grows in fertile soil and forest floors. Student A: How do people consume it for medicine? Student B: People boil them in water and drink the warm herbal water to soothe a sore throat, relieve coughs, or ease an upset stomach.`
	sirihModel := `Sirih leaves are a traditional medicinal plant. People boil them in water and drink the warm herbal water to soothe a sore throat, relieve coughs, or ease an upset stomach.`

	mcq := func(stem, a, b, c, d, correct string) seedTask {
		return quizTask(model.TaskMultipleChoice,
			"Choose the correct answer (A, B, C, or D), then say the complete sentence aloud: "+stem,
			`{"A":"`+a+`","B":"`+b+`","C":"`+c+`","D":"`+d+`"}`,
			`{"correct":"`+correct+`"}`)
	}

	return seedChapter{
		title: "Medicinal Plants",
		desc:  "Mendeskripsikan ciri, habitat, dan manfaat tanaman obat Melayu Riau dalam teks lisan deskriptif berbahasa Inggris.",
		courses: []seedCourse{
			{title: "Goal Setting", desc: "Tujuan pembelajaran yang hendak dicapai.", modules: []seedModule{
				{title: "My Learning Goals", typ: model.ModuleMonologue,
					content:    "By the end of the lesson: I can describe the kinds of medical plants in my place in English for 5-10 minutes correctly and fluently. I can explain the benefits of medical plants for ones to life healthy and wealthy in English for 5-10 minutes correctly and fluently.",
					transcript: "I can describe the kinds of medical plants in my place in English for 5 to 10 minutes correctly and fluently. I can explain the benefits of medical plants for ones to life healthy and wealthy in English correctly and fluently.",
					order:      1},
			}},
			{title: "Input & Exploration", desc: "Passages, vocabulary and pronunciation models.", modules: []seedModule{
				{title: "Living Pharmacy", typ: model.ModuleMonologue, content: "Read the passage about Riau Malay medicinal plants.", transcript: living, order: 1, tasks: []seedTask{
					speakTask("Read the passage aloud with clear pronunciation.", "pegagan => peh-gah-gahn; akar tunjuk langit => ah-kahr toon-jook lah-ngit; sirih => see-reeh; poultice => pohl-tiss"),
				}},
				{title: "Kinds of Medical Plants", typ: model.ModuleMonologue, content: "Fourteen plants from Table 2.1: Pandan, Nipah, Jeruju, Api-api, Bakau, Biduri, Pegagan, Mengkudu, Limau Purut, Akar pinang, Beluntas, Sirih, Akar tunjuk langit, Kupang-kupang.", transcript: plants, order: 2, tasks: []seedTask{
					speakTask("Name five plants from Table 2.1 and say what each treats.", "pandan => pahn-dahn; nipah => nee-pah; jeruju => jeh-roo-joo; bakau => bah-kow; biduri => bee-doo-ree; pegagan => peh-gah-gahn; mengkudu => meng-koo-doo; beluntas => beh-loon-tahs; sirih => see-reeh; kupang => koo-pahng"),
				}},
				{title: "Core Words", typ: model.ModuleMonologue, content: "Ten words with meanings and examples: root, sap, wound, tonic, heal, reduce, relieve, fragrant, antiseptic, bumpy.", transcript: coreWords, order: 3, tasks: []seedTask{
					speakTask("Read each word aloud, then use three in your own sentences.", "sap => sahp; wound => woond; tonic => tohn-ik; relieve => ree-leev; fragrant => fray-grahnt; antiseptic => an-tee-sep-tik; bumpy => buhm-pee"),
				}},
				{title: "Language Function", typ: model.ModuleMonologue, content: langFunc, order: 4},
				{title: "Grammar Focus", typ: model.ModuleMonologue, content: grammar, order: 5},
				{title: "Language Expression", typ: model.ModuleMonologue, content: langExpr, order: 6},
			}},
			{title: "Comprehension Check", desc: "Video, dialogues, quizzes and oral questions.", modules: []seedModule{
				{title: "Akar Tunjuk Langit (Video)", typ: model.ModuleVideo, content: "Watch the video about Akar Tunjuk Langit, then retell what you learned.", media: "https://drive.google.com/uc?export=download&id=1PjT0WawX6vHngDf4Tqke1B9ckFCZ49BU", order: 1, tasks: []seedTask{
					// ponytail: book lists 5 video questions with no keys — prompt-only OPEN_ENDED until supervisor supplies keys.
					quizTask(model.TaskOpenEnded, "After watching the video: what plant is it and where does it grow?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "After watching the video: which part is used as medicine?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "After watching the video: how do people prepare it?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "After watching the video: what are its health benefits?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "After watching the video: why is it important for Riau?", "", `{"answers":[]}`),
					speakTask("Retell the video content aloud in your own words.", "akar tunjuk langit => ah-kahr toon-jook lah-ngit"),
				}},
				{title: "Putri and Pak Cik Rahim", typ: model.ModuleDialogue, content: "Listen to the dialogue and practice it with a partner.", transcript: putri, order: 2, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", "akar tunjuk langit => ah-kahr toon-jook lah-ngit; veranda => veh-rahn-dah; vitality => vai-tah-lih-tee"),
				}},
				{title: "Plant Vocabulary Quiz", typ: model.ModuleQuiz, content: "Choose the correct answer, then say the complete sentence aloud.", order: 3, tasks: []seedTask{
					mcq("The most valuable part of Akar Tunjuk Langit is hidden underground. People boil its _______.", "Flower", "Root", "Fruit", "Leaf", "B"),
					mcq("The white _______ of the Biduri plant can cause skin irritation.", "Bark", "Stem", "Sap", "Seed", "C"),
					mcq("People apply crushed betel leaves directly to an open _______ to prevent infection.", "Wound", "Breath", "Stomach", "Cough", "A"),
					mcq("Villagers drink a warm herbal _______ to restore energy and stamina.", "Poison", "Tonic", "Flavor", "Spice", "B"),
					mcq("Pegagan contains compounds that help _______ minor skin cuts faster.", "Break", "Heal", "Harm", "Burn", "B"),
					mcq("Drinking Mengkudu juice is known to help _______ high blood pressure.", "Reduce", "Increase", "Raise", "Climb", "A"),
					mcq("Boiled pandan leaves can help _______ joint pain and body stiffness.", "Damage", "Worsen", "Relieve", "Create", "C"),
					mcq("Pandan leaves are used because of their sweet and _______ smell.", "Fragrant", "Bitter", "Rotten", "Sour", "A"),
					mcq("Sirih is popular as mouthwash because it has strong _______ properties.", "Toxic", "Allergic", "Antiseptic", "Synthetic", "C"),
					mcq("You recognize Limau Purut by its rough and _______ skin.", "Smooth", "Flat", "Bumpy", "Soft", "C"),
				}},
				{title: "Descriptive Text", typ: model.ModuleMonologue, content: "Read the descriptive text, then answer orally in complete sentences. Useful expressions: In my opinion... I think... It is important because... People traditionally use... We should preserve...", transcript: descriptive, order: 4, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "What is the scientific name of Tunjuk Langit?", "", `{"answers":["Helminthostachys zeylanica"]}`),
					quizTask(model.TaskOpenEnded, "Where does Tunjuk Langit usually grow?", "", `{"answers":["Moist and shady places near forests and rivers"]}`),
					quizTask(model.TaskOpenEnded, "Why is the plant called Tunjuk Langit?", "", `{"answers":["Its upright part looks like a small stick pointing to the sky"]}`),
					quizTask(model.TaskOpenEnded, "What part of the plant has traditionally been used in herbal preparations?", "", `{"answers":["The roots"]}`),
					quizTask(model.TaskOpenEnded, "Why is Tunjuk Langit important to Malay communities?", "", `{"answers":["It is part of Riau's local knowledge and natural heritage"]}`),
					speakTask("Answer in your own words: how can young people help preserve this knowledge?", "heritage => heh-rih-tij; preserve => preh-zerv"),
				}},
				{title: "Listening: Tunjuk Langit", typ: model.ModuleMonologue, content: "Match each beginning (1-5) with its ending (A-E), then pronounce your answers aloud. Source monologue: the Tunjuk Langit description.", transcript: matchText, order: 5, tasks: []seedTask{
					quizTask(model.TaskMatching, "Match beginnings 1-5 with endings A-E.",
						`{"statements":{"1":"The plant is called Tunjuk Langit because...","2":"Unlike open sunny fields, this plant naturally grows where...","3":"The most useful part taken from this fern is...","4":"Local people boil the root to make an herbal tonic that...","5":"The natural properties of this plant also..."},"endings":{"A":"...it thrives in damp forest floors and freshwater swamp edges.","B":"...help relieve backache, stiff joints, and body soreness.","C":"...its upright spike stands straight toward the sky.","D":"...its underground root system.","E":"...boosts stamina and restores physical energy."}}`,
						`{"1":"C","2":"A","3":"D","4":"E","5":"B"}`),
				}},
				{title: "Akar Oral Questions", typ: model.ModuleMonologue, content: "Answer each question orally with clear pronunciation.", transcript: akarMono, audio: "https://drive.google.com/uc?export=download&id=10y9H9iJm6KaX-WDVYGigU_nA85OifHfS", order: 6, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "What is the topic of the speaker's monologue?", "", `{"answers":["Akar Tunjuk Langit"]}`),
					quizTask(model.TaskOpenEnded, "What is the alternative name for Akar Tunjuk Langit?", "", `{"answers":["Kamraj"]}`),
					quizTask(model.TaskOpenEnded, "What type of plant is Akar Tunjuk Langit?", "", `{"answers":["Primitive terrestrial fern"]}`),
					quizTask(model.TaskOpenEnded, "Why did local people name this plant Tunjuk Langit?", "", `{"answers":["Its upright stem points straight toward the sky"]}`),
					quizTask(model.TaskOpenEnded, "What color are the leaves of this plant?", "", `{"answers":["Dark green"]}`),
					quizTask(model.TaskOpenEnded, "What kind of environment does this plant need?", "", `{"answers":["Moist, shaded tropical environment"]}`),
					quizTask(model.TaskOpenEnded, "Mention two habitats where this plant can be found.", "", `{"answers":["Tropical forest floors and freshwater swamp edges"]}`),
					quizTask(model.TaskOpenEnded, "Which part is most valuable for traditional healing?", "", `{"answers":["Its root"]}`),
					quizTask(model.TaskOpenEnded, "How is the root prepared before consuming?", "", `{"answers":["The root is boiled to create a restorative tonic"]}`),
					quizTask(model.TaskOpenEnded, "What are the primary health benefits of the tonic?", "", `{"answers":["Boosts body stamina and energy"]}`),
				}},
				{title: "Maya and Nek Minah", typ: model.ModuleDialogue, content: "Fill in the blanks using the word bank, then practice speaking the dialogue aloud.", transcript: mayaGap, order: 7, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Fill blanks (1-8) from: Boil, Wound, Heal, Fragrant, Antiseptic, Relieve, Leaves, Reduce. Then read the complete dialogue aloud.",
						`{"bank":["Boil","Wound","Heal","Fragrant","Antiseptic","Relieve","Leaves","Reduce"]}`,
						`{"blanks":["Leaves","Fragrant","Boil","Reduce","Antiseptic","Wound","Heal","Relieve"]}`),
				}},
			}},
			{title: "Guided Speaking", desc: "Pegagan and Sirih guided questions and vocabulary.", modules: []seedModule{
				{title: "Pegagan and Sirih Guide", typ: model.ModuleMonologue, content: "Guiding questions: What is the scientific name of Pegagan? Of Sirih leaves? What does it look like and where does it grow? Which part is used and how is it prepared? What are the health benefits? Sentence starters: Pegagan is a traditional medicinal plant from... It usually grows in... You can recognize it by its... People usually use the [leaves / root / bark] to... To consume it, they simply... It is very useful because it can help [heal / relieve / reduce]... Vocabulary: " + guideVocab, order: 1, tasks: []seedTask{
					speakTask("Answer the guiding questions orally about Pegagan or Sirih.", "pegagan => peh-gah-gahn; sirih => see-reeh; boil => boyl; crush => krahsh; moist => moyst"),
				}},
				{title: "Sirih Dialogue Model", typ: model.ModuleDialogue, content: "Practice the example dialogue with a partner, then record your voice.", transcript: sirihDialog, order: 2, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", "sirih => see-reeh; soothe => sooth"),
				}},
			}},
			{title: "Independent Speaking", desc: "Present one medicinal plant without reading.", modules: []seedModule{
				{title: "Plant Presentation", typ: model.ModuleOralTest, content: "Choose one Riau medical plant. Speak for 3 to 5 minutes without reading: physical appearance, preparation or planting, and benefits. Language functions: Describing (Pegagan is a traditional medical plant). Asking (Where is it planted?). Giving information (What is its specification?). Opinions and reasons (I think pegagan plants can make my body fit because it is healthy).", transcript: sirihModel, order: 1, tasks: []seedTask{
					speakTask("Present your chosen plant for 3 to 5 minutes without reading, then record.", "sirih => see-reeh; pegagan => peh-gah-gahn"),
				}},
				{title: "Language Functions Reference", typ: model.ModuleMonologue, content: "Describing: Pagagan is a traditional medical plant. Asking: Where is it planted? Giving information: What is its specification? Activity: How to make a healthy drink? Experience: I planted it last year. Likes and opinions: I like the healthy drink. I think pagagan plants can make my body fit. Reasons: I like it because it is healthy. Wishes, agreeing, responding, appreciation frames. Student model: Tooth paste is made from sirih leaves. Sirih leaves eaten people to maintain teeth strong.", order: 2},
			}},
		},
	}
}
