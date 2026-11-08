// Chapter 1 rebuilt from Seeder Brief Chapter 1 only (supervisor call).
// 7 courses = GECISER phases; every module maps to a brief item; keys live
// in AnswerKeyJSON and reveal after an attempt, except the Oral Questions
// whose answers show outright. ASR Feedback / Self-Reflection / Mastery are
// task-less modules (frontend gates) — no schema change by design.
package main

import "asri-backend/internal/model"

func ch1Final() seedChapter {
	history := `Riau coastal traditional games have been part of Malay community life for many generations. They were usually played by children, teenagers, and adults during their free time or at community gatherings. Many of these games used simple materials from the local environment, such as wood, coconut shells, bamboo, rattan, seeds, and shells. This shows the close relationship between the Malay people and their natural environment. Traditional games were not only for fun. They also helped people develop important skills and values. Children learned how to cooperate, communicate, follow rules, concentrate, solve problems, and respect other players. One example is Rimau, a traditional strategy game associated with Malay coastal villages. It was traditionally played by teenagers and adults, especially in fishing communities. Historical information from the Riau Provincial Government states that Rimau was popular around the 1890s, although there is no exact record of when it first began. The game was passed down from one generation to another. Other traditional games associated with Riau include gasing, congkak, bakiak, layang-layang, and sengki. These games have different rules and equipment, but they share an important function: they bring people together and help preserve cultural traditions. Today, some traditional games are becoming less common because of changes in lifestyle and the popularity of digital entertainment. Therefore, schools, families, and cultural organizations can help introduce these games to young people. Riau coastal traditional games are valuable cultural heritage. By learning and playing them, young people can enjoy themselves while learning about the history, environment, and values of Riau Malay culture.`
	kinds := `Bakiak => A traditional team game where players wear long wooden sandals and walk together in the same direction. ;; Papan congkak => A traditional board game where players move small seeds or beads around holes on a board. ;; Selop => A traditional game using slippers, usually played by arranging, throwing, or balancing slippers according to the rules. ;; Gasing => A traditional spinning-top game where players spin a top and try to keep it spinning. ;; Gasing pangkah => A spinning-top game where players spin their tops and try to hit or strike another player's top. ;; Gasing uri => A spinning-top game that focuses on how long a top can spin. ;; Layang-layang => A traditional kite game where players fly kites in the air using a string. ;; Jong layo => A traditional Riau game using miniature boats that are placed on water and moved by wind or water. ;; Pacu kolek => A traditional Riau boat-racing game where participants paddle boats and race to the finish line. ;; Lulu Cina Buta => A traditional group game where one blindfolded player tries to catch the other players. ;; Ular naga => A traditional group game where players form a line and pass through a human gate. ;; Selambut => A traditional game played with small objects, usually involving throwing, catching, and picking them up according to certain rules. ;; Benteng => A traditional team game where players try to capture the opposing team's base while protecting their own base. ;; Boi Boian / pecah piring => A traditional team game where players knock down a stack of objects and try to rebuild it while avoiding being hit by a ball. ;; Engrang => A traditional game where players walk or race using tall bamboo stilts. ;; Statak / tengkak => A traditional hopping game where players jump on one foot through marked squares on the ground. ;; Galah Panjang => A team game where attackers try to run through a grid of guarding defenders and back without getting tagged.`
	jongShort := `Jong Layo is a traditional boat game from Riau, Indonesia. The game uses miniature boats that are usually made and decorated creatively. The boats are placed on the water and can move with the help of wind and water. Jong Layo is not only a fun activity but also reflects the creativity and cultural traditions of the people in Riau.`
	jongHowto := `Jong Layo, a traditional coastal game associated with Malay maritime culture. This game reflects the creativity and maritime traditions of Malay communities in Riau. Jong Layo is played using a small traditional sailing boat. The boat is designed to move across the water with the help of the wind. It is an interesting game that teaches us about the sea, wind, and traditional boat-making skills. First, we need to prepare a small sailing boat called a jong. The boat usually has a hull, a mast, and a sail. We also need a suitable coastal area with enough wind and relatively calm water. Second, we check the boat to make sure that its sail and mast are properly attached. We must ensure that the boat is ready to sail. Third, we place the boat on the water and let the wind fill its sail. The wind pushes the sail and moves the boat forward. We watch carefully as the boat travels across the water. Next, we observe the movement of the boat. If we are playing with friends, we can race our boats and see which one travels the fastest or reaches the finish line first, depending on the rules of the game. Finally, we collect our boats and keep the playing area clean. We must also be careful when playing near the sea and follow safety instructions. In conclusion, Jong Layo is more than just a traditional game. It introduces us to Malay coastal culture, develops our creativity, and helps us appreciate the maritime heritage of Riau. We should preserve this traditional game so that future generations can learn about and enjoy it.`
	boatParts := `Traditional sailing boat => Perahu layar tradisional. ;; Hull => Badan perahu. ;; Mast => Tiang layar. ;; Sail => Layar. ;; Wind => Angin. ;; Coast => Pesisir pantai. ;; Finish line => Garis finis. ;; Maritime heritage => Warisan maritim.`
	gameWords := `Miniature boat => A small model of a boat. ;; Boat => A vehicle used to travel on water. ;; Water => A liquid that covers rivers, lakes, and seas. ;; Competition => An event where people compete against each other. ;; Tradition => The customs, traditions, and way of life of a group of people. ;; Creativity => The ability to create new and original ideas. ;; Game => The collective arena of competition, strategy, and continuous play where teams test their skills, adapt to challenges, and strive to improve. ;; Culture => The shared values, habits, and working style that shape a team's identity. ;; Team => A group of people who work together, support one another, and combine their strengths to achieve a common goal. ;; Move => Taking proactive action, showing agility, and moving forward without getting stuck in hesitation. ;; Win => The achievement, success, or milestone reached after hard work and collaboration. ;; Lose => A setback or failure that serves as a valuable lesson to grow, adapt, and come back stronger instead of giving up.`
	nouns := `Traditional Game Culture Local Player Sail Team Decorate Float Watch Group Partner Scoring board Playground Ball Rope Stick Stone Seed Board Line Circle Square`
	verbs := `Play Run Jump Hope Throw Catch Kick Spin Chase Hide Touch Move Count Win Lose Seek`
	adjs := `Fun Exciting Interesting Challenging Easy Difficult Enjoyable Traditional Active Competitive Creative Educational`
	langFunc := `Asking information: What game is this? What is the game called? How do you play it? How many players do you need? What equipment do you need? Where can we play it? Who can play this game? What are the rules? Giving information: It is called Congklak. It is a traditional game. Two players can play the game. We need a board and small seeds. Players sit opposite each other. The game is usually played indoors. Explaining how to play: First, prepare the board and seeds. Next, put the seeds into the holes. Then, choose a player to start. After that, move the seeds. Finally, count the seeds and decide the winner. Likes: I like Congklak. I really like Gasing. I enjoy playing Galah Panjang. My favorite game is Congklak. Dislikes: I don't like this game. I don't really like difficult games. Reasons: I like Galah Panjang because it is exciting and I can play it with friends. Inviting: Do you want to play Congklak? Would you like to play Gasing? Let's play Galah Panjang! Why don't we play Ular Naga? Would you like to join us? Accepting: Sure! Yes, I'd love to. That sounds fun! Okay, let's play. Refusing politely: Sorry, I can't. Maybe later. I'm sorry, but I don't have time.`
	grammar := `Simple present: Two players play Congklak. Players move the seeds. The game needs two teams. Nominal sentence: Jong Layo is a unique traditional game from Riau. Verbal sentence: The players put the small boat on the water. Imperatives: Stand in a line. Choose a player. Throw the ball. Run to the line. Count the points. Modals: You can play this game with your friends. Two people can play Congklak. Sequencing: First, Next, Then, After that, Finally.`
	langExpr := `Make => Players make and decorate their own small boats creatively. (Para pemain membuat dan menghias perahu kecil mereka sendiri secara kreatif.) ;; Sail => The wooden boats sail smoothly when the wind blows. (Perahu-perahu kayu itu berlayar dengan lancar ketika angin bertiup.) ;; Watch => Children watch the boats float across the pond. (Anak-anak menonton perahu-perahu itu mengapung di atas kolam.) ;; Combine => The game combines creativity, tradition, and fun. (Permainan ini menggabungkan kreativitas, tradisi, dan keseronokan.) ;; Prepare => Participants prepare their boats before the race starts. (Peserta mempersiapkan perahu mereka sebelum perlombaan dimulai.) ;; Traditional => Jong Layo is a popular traditional game in coastal regions of Riau. (Jong Layo adalah permainan tradisional yang populer di wilayah pesisir Riau.) ;; Creative => Each player makes a creative design for their wooden boat. (Setiap pemain membuat desain yang kreatif untuk perahu kayu mereka.) ;; Unique => The game is unique because it relies entirely on natural wind and water. (Permainan ini unik karena sangat mengandalkan angin dan air alami.) ;; Fun => Playing Jong Layo together is an exciting and fun outdoor activity. (Bermain Jong Layo bersama-sama adalah aktivitas luar ruangan yang menyenangkan dan mengasyikkan.) ;; Small => The game uses a small wooden boat that floats easily. (Permainan ini menggunakan perahu kayu kecil yang mengapung dengan mudah.)`
	paiman := `Juminah: Hi! What are you doing? Paiman: Hi! I'm learning about a traditional game from Riau called Jong Layo. Juminah: Jong Layo? What is it? Paiman: Jong Layo is a traditional game that uses a small boat. Juminah: How do you play it? Paiman: First, we prepare and decorate a small boat. Then, we put the boat on the water. Juminah: Does the boat move by itself? Paiman: The boat can move with the wind or water. Juminah: That sounds interesting. Is it difficult to play? Paiman: No, I think it is quite simple and fun. Juminah: What can we learn from this game? Paiman: We can learn creativity and appreciate the traditional culture of Riau. Juminah: I think I want to try it! Paiman: Sure! It would be fun to try it together.`
	jongLong := `Hello everyone. Today, I want to talk about a traditional game from Riau called Jong Layo. Jong Layo is a traditional game that uses a small boat. The boat is usually made and decorated creatively. Players put the boat on the water and make it move or sail. The game can be played by several people. The players prepare their boats and then put them on the water. They can watch the boats float and move with the wind or water. Jong Layo is an interesting traditional game because it is not only fun but also creative. Players can learn how to make and decorate a small boat. It also helps us appreciate the traditional culture of Riau. I think Jong Layo is a unique game because it combines creativity, tradition, and fun.`
	sarahDaveGap := `Sarah: Hi, Dave! What are you doing? Dave: Hi, Sarah! I'm learning about a traditional game from (1) ________ called Jong Layo. Sarah: Oh, what is Jong Layo? Dave: It is a traditional game that uses a (2) ________. Sarah: How do people play it? Dave: First, the (3) ________ prepare their boats. They also (4) ________ the boats. Sarah: Where do they put the boats? Dave: They put the boats on the (5) ________. Sarah: How do the boats move? Dave: They can move with the (6) ________ or water. Sarah: That sounds interesting! Is the game difficult? Dave: No. It is fun and (7) ________. Sarah: What can people learn from Jong Layo? Dave: They can learn about creativity and appreciate the traditional (8) ________ of Riau.`
	galah := `Galah Panjang is a traditional team game. Players are divided into two teams. One team tries to cross the lines while the other team guards the lines. Players need to run and move quickly. The players cannot be touched by the guards. The game needs teamwork and strategy. It is fun and exciting to play with friends. Galah Panjang can be played by adults or children.`
	tengku := `Tengku: Hey, we have some free time tomorrow afternoon. Do you want to hang out with the kids from the neighborhood? Siti: Sounds good, I'm getting bored of staying home and staring at my phone all day. What do you have in mind? Tengku: How about we get them to play Galah Panjang at the field near the village hall? It's been ages since we last played that traditional game. Siti: Oh, that's a great idea! But it usually requires quite a few people for two teams. Do we have enough players? Tengku: Don't worry, I chatted with a few of them briefly yesterday, and they're totally down. We can split into two teams, with at least three or four people each. Siti: Awesome! But get ready to run a lot—the line defenders have to be super agile so the opposing team can't slip past to the finish line. Tengku: Exactly, it's a great afternoon workout so our bodies aren't so stiff. I'll take care of bringing the chalk to draw the field lines. Siti: Alright then! I'll help message the others in the WhatsApp group so everyone heads to the field tomorrow afternoon.`
	galahVocab := `Miniature boat => A small model of a boat. The player launches a lightweight miniature boat into the sea. ;; Boat => A vehicle used to travel on water. This unique sailing boat moves entirely powered by the wind. ;; Water => A liquid that covers rivers, lakes, and seas. The tiny vessel glides smoothly over the open water. ;; Competition => An event where people compete against each other. Local fishermen cheer loudly during the annual Jong competition. ;; Tradition => A custom, tradition, and way of life of a group of people. Racing Jong is a cherished coastal tradition in Riau. ;; Creativity => The ability to create new and original ideas. Crafting a winning Jong demands high artistic creativity. ;; Traditional boat race => A competitive sporting event where teams race traditional regional boats by rowing together. The annual traditional boat race attracts many visitors to the riverbank. ;; Kolek => A traditional wooden canoe used by Malay communities for traveling and fishing. Fishermen use a kolek to navigate through narrow rivers. ;; Race => A contest where participants compete to see who is the fastest to finish. Every team paddled as fast as they could to win the race. ;; Team => A group of people working or playing together to achieve a common goal. Each team consists of ten strong paddlers. ;; Culture => The shared customs, arts, and traditions passed down through generations in a society. We must preserve our local culture so it will never fade away. ;; Cooperation => The act of working together harmoniously to achieve a shared objective. Good cooperation helps the team move the boat smoothly. ;; Local community => A group of people sharing the same geographical area and cultural background. The local community enthusiastically supports the traditional festival. ;; Prize => A reward or trophy given to the winner of a competition. The champion received a trophy and a cash prize. ;; Winner => The person or group that finishes first or wins a competition. The committee announced village A as the winner. ;; River => A large natural stream of fresh water flowing across the land. People gathered along the wide river to watch the boats. ;; Water => The clear liquid flowing in rivers or oceans that allows boats to float. The wooden boat glided smoothly over the calm water.`
	pacu := `Pacu Kolek is made from a coconut tree. Pacu Kolek is played in a team that consists of 4 to 10 players.`

	mcq := func(n int, stem, a, b, c, d, correct string) seedTask {
		return quizTask(model.TaskMultipleChoice,
			"Choose the correct answer (A, B, C, or D), then say the complete sentence aloud: "+stem,
			`{"A":"`+a+`","B":"`+b+`","C":"`+c+`","D":"`+d+`"}`,
			`{"correct":"`+correct+`"}`)
	}

	return seedChapter{
		title: "Traditional Games",
		desc:  "Mendeskripsikan dan menjelaskan aturan permainan tradisional Melayu Riau pesisir secara lisan dalam bahasa Inggris.",
		courses: []seedCourse{
			{title: "1. Goal Setting", desc: "Goal Setting — kemampuan yang dicapai setelah mempelajari topik ini (brief #1).", modules: []seedModule{
				{title: "What You Will Learn", typ: model.ModuleMonologue,
					content:    "By the end of the lesson: I can explain Jong Layo in English for 5-7 minutes correctly and fluently; I can describe how to play Jong Layo in English for 5-7 minutes correctly and fluently. After learning this topic, you are able to:",
					transcript: "After learning this topic, you are able to: name several Riau Malay traditional games; describe a traditional game using simple English; explain how to play a traditional game; ask and answer questions about traditional games; express likes and dislikes; give simple opinions about traditional games; invite friends to play a traditional game; demonstrate and present a traditional game in English.",
					order:      1, tasks: []seedTask{
						quizTask(model.TaskOpenEnded, "After learning this topic, you are able to:",
							`{"objectives":["Name several Riau Malay traditional games","Describe a traditional game using simple English","Explain how to play a traditional game","Ask and answer questions about traditional games","Express likes and dislikes","Give simple opinions about traditional games","Invite friends to play a traditional game","Demonstrate and present a traditional game in English"]}`,
							``),
					}},
			}},
			{title: "2. Input & Exploration", desc: "Input & Exploration — monologue, Table 4, Jong Layo, vocabulary, pronunciation, expressions, Congklak video (brief #2-13).", modules: []seedModule{
				{title: "The History Monologue", typ: model.ModuleMonologue,
					content:    "Please listen to the monologue carefully. Repeat it clearly and loudly.",
					transcript: history, order: 1, tasks: []seedTask{
						speakTask("Listen to the monologue, then repeat it clearly and loudly.", ""),
						seedTask{typ: model.TaskSpeakingRecording, prompt: "What are Riau coastal traditional games?", options: `{"Model answer":"Riau coastal traditional games are traditional games played by Malay communities in the coastal areas of Riau."}`},
						seedTask{typ: model.TaskSpeakingRecording, prompt: "Mention three traditional games from coastal Riau.", options: `{"Model answer":"Three traditional games are gasing, bakiak, and congkak."}`},
						seedTask{typ: model.TaskSpeakingRecording, prompt: "Why are traditional games important for Riau Malay culture? Give at least two reasons.", options: `{"Model answer":"Traditional games are important because they help preserve Riau Malay cultural heritage. They also bring people together and teach values such as teamwork, patience, discipline, and respect for others."}`},
						seedTask{typ: model.TaskSpeakingRecording, prompt: "What materials are commonly used to make traditional game equipment in coastal Riau?", options: `{"Model answer":"Traditional games commonly use wood, coconut shells, bamboo, rattan, seeds, and shells."}`},
						seedTask{typ: model.TaskSpeakingRecording, prompt: "What can children learn from playing traditional games such as gasing, bakiak, or congkak? Give at least three values.", options: `{"Model answer":"Children can learn many important values from traditional games. They can learn teamwork, communication, patience, concentration, discipline, coordination, and fair play and appreciate their local culture."}`},
					}},

				{title: "ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the ASR feedback to know your speaking level. When you reach 80% correct, please continue to the next step. My score is: ___%.",
					order:   2},

				{title: "Table 4: Kinds of Games", typ: model.ModuleMonologue,
					content:    "Table 4 RIAU MALAY TRADITIONAL GAMES: games played on the water and on land, with meanings.",
					transcript: kinds, order: 1, tasks: []seedTask{
						speakTask("Name five games from Table 4 and say what each uses.", ""),
					}},

				{title: "Jong Layo", typ: model.ModuleMonologue,
					content:    "Listen to the monologue, then read it aloud.",
					transcript: jongShort, order: 1, tasks: []seedTask{
						speakTask("Read the Jong Layo monologue aloud.", ""),
					}},
				{title: "How to Play", typ: model.ModuleMonologue,
					content:    "Children and adults can play Riau Malay traditional games to celebrate national independence days, regional independence days, cultural ceremonies, and other activities. Please listen to the monologue about how to play Jong Layo. Repeat loudly and clearly.",
					transcript: jongHowto, order: 2, tasks: []seedTask{
						speakTask("Read the How to Play Jong Layo monologue aloud.", ""),
						quizTask(model.TaskOpenEnded, "What materials or parts are needed to prepare a jong?", "",
							`{"answers":["We need a small traditional sailing boat called a jong. It has a hull, a mast, and a sail. We also need a suitable coastal area with enough wind and relatively calm water"]}`),
						quizTask(model.TaskOpenEnded, "How does the wind move the boat?", "",
							`{"answers":["The wind fills the sail and pushes the boat forward across the water"]}`),
						quizTask(model.TaskOpenEnded, "What should we check before putting the boat on the water?", "",
							`{"answers":["We should check the hull, mast, and sail to make sure they are properly attached and the boat is ready to sail"]}`),
						quizTask(model.TaskOpenEnded, "How can we preserve Jong Layo as part of Malay culture?", "",
							`{"answers":["We can preserve Jong Layo by learning how to play it, practicing it with our friends, participating in traditional games and cultural events, and teaching younger generations about this traditional game"]}`),
					}},

				{title: "Boat Parts", typ: model.ModuleMonologue,
					content:    "Words related to the Jong Layo game, with Indonesian meanings.",
					transcript: boatParts, order: 1},
				{title: "Game Words", typ: model.ModuleMonologue,
					content:    "Twelve words with meanings.",
					transcript: gameWords, order: 2},

				{title: "Nouns", typ: model.ModuleMonologue,
					content:    "Listen and repeat each noun.",
					transcript: nouns, order: 1},
				{title: "Verbs", typ: model.ModuleMonologue,
					content:    "Listen and repeat each verb.",
					transcript: verbs, order: 2},
				{title: "Adjectives", typ: model.ModuleMonologue,
					content:    "Listen and repeat each adjective.",
					transcript: adjs, order: 3},

				{title: "Expressions", typ: model.ModuleMonologue, content: langExpr, order: 1, tasks: []seedTask{
					speakTask("How many teams are there?", ""),
					speakTask("What does one team do?", ""),
					speakTask("What does the other team do?", ""),
					speakTask("What skills do players need?", ""),
					speakTask("Why is the game fun?", ""),
				}},

				{title: "ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the ASR feedback to know your speaking level. When you reach 80% correct, please continue to the next step. My score is: ___%.",
					order:   11},

				{title: "Congklak Games (Video)", typ: model.ModuleVideo,
					content: "The Congklak game is one of the Riau traditional games. It is usually played by kids and women. Please watch the video carefully.",
					media:   "https://www.youtube.com/watch?v=An0syZQapfA", order: 1},
				{title: "Language Function", typ: model.ModuleMonologue, content: langFunc, order: 2},
				{title: "Grammar Focus", typ: model.ModuleMonologue, content: grammar, order: 3},

				{title: "ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the ASR feedback to know your speaking level. When you reach 80% correct, please continue to the next step. My score is: ___%.",
					order:   15},
			}},
			{title: "3. Comprehension Check", desc: "Comprehension Check — video, dialogues, quizzes and oral questions (brief #14-20).", modules: []seedModule{
				{title: "How to Make Jong Layo (Video)", typ: model.ModuleVideo, content: "Watch how to make Jong Layo and how to play it.", media: "https://drive.google.com/uc?export=download&id=1QvjMecLLTMBoWvuDhK0S6paChAQUZwQo", order: 1, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Jong Layo is commonly played by children in _______.", "", `{"answers":["Riau"]}`),
					quizTask(model.TaskOpenEnded, "What materials are used to make Jong Layo?", "", `{"answers":["Tree trunk, plastic, and rope"]}`),
					quizTask(model.TaskOpenEnded, "How is the procedure to play Jong Layo?", "", `{"answers":["Players release their mini boats into the water from a starting line, letting the wind drive them freely, and the boat that reaches the finish line first wins"]}`),
				}},
				{title: "Paiman and Juminah", typ: model.ModuleDialogue, content: "Listen to the dialogue and answer the questions.", transcript: paiman, order: 2, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", ""),
				}},
				{title: "Jong Layo Vocabulary Quiz", typ: model.ModuleQuiz, content: "Choose the correct answer, then say the complete sentence aloud.", order: 3, tasks: []seedTask{
					mcq(1, "Jong Layo is a traditional ________ from Riau.", "Game", "Food", "Song", "Car", "A"),
					mcq(2, "Jong Layo is a traditional game that uses a small ________.", "Kite", "Boat", "Ball", "Drum", "B"),
					mcq(3, "First, we ________ and decorate a small boat.", "Prepare", "Catch", "Throw", "Push", "A"),
					mcq(4, "Then, we put the boat on the ________.", "Road", "Field", "Water", "Wall", "C"),
					mcq(5, "The boat can move with the ________ or water.", "Wind", "Player", "Song", "Rope", "A"),
					mcq(6, "No, I think it is quite simple and ________.", "Difficult", "Dangerous", "Fun", "Easy", "C"),
					mcq(7, "We can learn ________ from this game.", "Creativity", "Swimming", "Cooking", "Cheating", "A"),
					mcq(8, "Jong Layo is part of the traditional ________ of Riau.", "Culture", "Sport", "Food", "Drink", "A"),
					mcq(9, "The players move ________ when they play the game.", "Together", "Boat", "Traditional", "Local", "A"),
					mcq(10, "Jong Layo ________ a traditional game from Riau.", "Prepare", "Is", "Decorate", "Are", "B"),
				}},
				{title: "Sarah and Dave", typ: model.ModuleDialogue, content: "Fill in the blanks using the word bank, then practice speaking the dialogue aloud.", transcript: sarahDaveGap, order: 4, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Fill blanks (1-8) from: Riau, Water, Wind, Culture, Small boat, Decorate, Creative, Players. Then read the complete dialogue aloud.",
						`{"bank":["Riau","Water","Wind","Culture","Small boat","Decorate","Creative","Players"]}`,
						`{"blanks":["Riau","Small boat","Players","Decorate","Water","Wind","Creative","Culture"]}`),
				}},
				{title: "Listening: Jong Layo", typ: model.ModuleMonologue, content: "Match each beginning (1-5) with its ending (A-E), then pronounce your answers aloud. Source monologue: the Jong Layo presentation.",
					transcript: jongLong, audio: "https://drive.google.com/uc?export=download&id=1xo2Ky9v42m1hFD5cINFGxUDnKuUJTG3R", order: 5, tasks: []seedTask{
						quizTask(model.TaskMatching, "Match beginnings 1-5 with endings A-E.",
							`{"statements":{"1":"Today, the speaker wants to talk about...","2":"Jong Layo is a traditional game that uses...","3":"Players put the small boat on the water to...","4":"Players can learn creativity by learning...","5":"The speaker thinks Jong Layo is unique because it can..."},"endings":{"A":"...how to make and decorate a small boat.","B":"...a traditional game from Riau called Jong Layo.","C":"...combine creativity, tradition, and fun.","D":"...make it move or sail.","E":"...a small boat that is decorated creatively."}}`,
							`{"1":"B","2":"E","3":"D","4":"A","5":"C"}`),
					}},
				{title: "Jong Layo Oral Questions", typ: model.ModuleMonologue, content: "Answer each question orally with clear pronunciation. Answers are shown.",
					transcript: jongLong, order: 6, tasks: []seedTask{
						quizTask(model.TaskOpenEnded, "What is the name of the traditional game from Riau?", `{"Model answer":"Jong Layo"}`, `{"answers":["Jong Layo"]}`),
						quizTask(model.TaskOpenEnded, "What does Jong Layo use?", `{"Model answer":"A small boat"}`, `{"answers":["A small boat"]}`),
						quizTask(model.TaskOpenEnded, "How is the boat usually made and decorated?", `{"Model answer":"Creatively"}`, `{"answers":["Creatively"]}`),
						quizTask(model.TaskOpenEnded, "Where do the players put the boat?", `{"Model answer":"Water"}`, `{"answers":["Water"]}`),
						quizTask(model.TaskOpenEnded, "How can the boat move?", `{"Model answer":"Wind or water"}`, `{"answers":["Wind or water"]}`),
						quizTask(model.TaskOpenEnded, "How many people can play the game?", `{"Model answer":"2 or more"}`, `{"answers":["2 or more"]}`),
						quizTask(model.TaskOpenEnded, "What can players learn from Jong Layo?", `{"Model answer":"How to make and decorate a small boat"}`, `{"answers":["How to make and decorate a small boat"]}`),
						quizTask(model.TaskOpenEnded, "Why is Jong Layo an interesting traditional game?", `{"Model answer":"Because it is a miniature sailing boat race and one of the traditional games of Riau"}`, `{"answers":["Because it is a miniature sailing boat race and one of the traditional games of Riau"]}`),
						quizTask(model.TaskOpenEnded, "What does Jong Layo combine?", `{"Model answer":"Creativity, tradition, and fun"}`, `{"answers":["creativity, tradition, and fun"]}`),
						quizTask(model.TaskOpenEnded, "What does Jong Layo help people appreciate?", `{"Model answer":"The traditional culture of Riau"}`, `{"answers":["The traditional culture of Riau"]}`),
					}},
				{title: "Story Order", typ: model.ModuleMonologue, content: "Rearrange the sentences into the correct order, then speak the complete story aloud.", transcript: "A. The players prepare their boats and put them on the water. B. Jong Layo is a traditional game from Riau. C. Players can learn how to make and decorate a small boat. D. The game can be played by several people. E. Jong Layo combines creativity, tradition, and fun. F. The boat is usually made and decorated creatively.", order: 7, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "Rearrange A-F into the correct story order.", `{"sentences":["A","B","C","D","E","F"]}`, `{"order":["B","F","D","A","C","E"]}`),
				}},

				{title: "ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the ASR feedback to know your speaking level. When you reach 80% correct, please continue to the next step. My score is: ___%.",
					order:   8},
			}},
			{title: "4. Guided Speaking", desc: "Guided Speaking — Galah Panjang monologue, dialogue and vocabulary (brief #21-24).", modules: []seedModule{
				{title: "Galah Panjang", typ: model.ModuleMonologue, content: "Listen to the monologue and answer the guiding questions orally: Have you ever played Galah Panjang? Do you know how to make Galah Panjang? What materials are used to make Galah? Where can you get the materials?", transcript: galah, order: 1, tasks: []seedTask{
					speakTask("Answer the guiding questions orally and record your answer.", ""),
					quizTask(model.TaskOpenEnded, "Complete with words from the text: Galah Panjang is a traditional (1)____ game. Players are divided into (2)____ teams. One team tries to (3)____ the lines. The game needs teamwork and (4)____. It can be played by (5)____ or children.",
						"", `{"blanks":["team","two","cross","strategy","adults"]}`),
				}},
				{title: "Tengku and Siti", typ: model.ModuleDialogue, content: "Practice the dialogue with your partner, then record your voice.", transcript: tengku, order: 2, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", ""),
				}},
				{title: "Galah Panjang Vocabulary", typ: model.ModuleMonologue, content: "Please listen to the pronunciation model and repeat.", transcript: galahVocab, order: 3},

				{title: "ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the ASR feedback to know your speaking level. When you reach 80% correct, please continue to the next step. My score is: ___%.",
					order:   4},
			}},
			{title: "5. Independent Speaking", desc: "Independent Speaking — present a traditional game without reading the script (brief #25-26).", modules: []seedModule{
				{title: "My Presentation", typ: model.ModuleOralTest, content: "Develop a monologue about Riau Malay Traditional Games. Speak for 3 to 5 minutes without reading the script. Describe the physical appearance, preparation, and the benefits. Choose one game: Bakiak, Gasing, Layang-layang, Pacu kolek, Lulu Cina Buta, Selambut, Benteng, Engrang. Record the monologue by clicking the record button.", transcript: pacu, order: 1, tasks: []seedTask{
					speakTask("Present your monologue for 3 to 5 minutes without reading, then record.", ""),
				}},
				{title: "Language Functions Reference", typ: model.ModuleMonologue, content: "Describing: Jong Layo is a unique traditional game from Riau. Asking: How do you play it? Giving information: Two players can play the game. Opinions: I think it is exciting. Reasons: I like Galah Panjang because it is exciting and I can play it with friends. Inviting: Let's play Galah Panjang!", order: 2},

				{title: "ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the ASR feedback to know your speaking level. When you reach 80% correct, please continue to Self-Reflection. My score is: ___%.",
					order:   3},
			}},
			{title: "6. Self-Reflection", desc: "Self-Reflection — evaluate speaking, identify difficulty, plan improvement (brief #27-28).", modules: []seedModule{
				{title: "My Reflection", typ: model.ModuleMonologue,
					content: "Did I speak fluently? Did I pronounce the words correctly? Did I use appropriate vocabulary? What was my biggest difficulty? What should I improve?",
					order:   1},
				{title: "Overall ASR Feedback", typ: model.ModuleMonologue,
					content: "Check the overall ASR feedback for this chapter. When you reach 80% correct, continue to Re-practice and Mastery. My score is: ___%.",
					order:   2},
			}},
			{title: "7. Re-practice and Mastery", desc: "Re-practice low scores, enrichment for high scores, then next lesson (brief #29).", modules: []seedModule{
				{title: "Practice and Mastery", typ: model.ModuleMonologue,
					content: "Speaking 1 > ASR Feedback > Self-Reflection > Re-practice (low scores). Speaking 2 > ASR Feedback > Self-Reflection > Enrichment (high scores). If the goal is reached: next lesson.",
					order:   1},
			}},
		},
	}
}
