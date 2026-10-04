// Chapter 4 content frozen from GIECSER_FINAL Lesson 4 (pp.64-82):
// goal setting → independent speaking. Self-reflection questions, the
// speaking rubric and the repractice loop live in the frontend — no rows
// by design. Unlike Chapters 1-3 the book gives no MCQ/matching keys here;
// comprehension is oral Q&A, seeded as OPEN_ENDED with the book's answers.
package main

import "asri-backend/internal/model"

func ch4Final() seedChapter {
	heritage := `Riau Malay traditional clothing is a graceful heritage that reflects Islamic values, modesty, and personal dignity. The attire is typically characterized by loose-fitting cuts that cover the body politely while allowing comfort in the tropical climate. In general, men wear a long tunic and matching trousers paired with a kain samping (a woven waistcloth) around the hips and a tanjak or black cap on the head. Women wear graceful garments like Baju Kurung or Baju Kebaya, complemented by a head covering (tudung) and hand-woven songket fabrics. Beyond their beauty and shimmering golden threads, these clothes are deeply respected because every collar, button, and fold symbolizes cultural pride, good manners, and moral character.`
	attireWords := `Tenun songket Tenun siak Tekat melayu Tanjak Selendang Baju cekak musang Kopiah songkok Kain samping Capal terompah Baju kurung teluk belanga Tudung manto Dokoh Pending sabuk Selop sandal`
	langFunc := `Asking about clothing: What is this? What is it called? What color is it? What is it made of? Who wears it? When do people wear it? Giving information: It is called Baju Melayu. It is a traditional Malay outfit. It is usually worn by men. It is made of fabric. It is usually worn during traditional ceremonies. Describing clothing: It is beautiful. It is colorful. It has a beautiful pattern. It is made of silk. It has long sleeves. It is usually worn with a songkok. Expressing opinion: I think it is beautiful. I think it looks elegant. In my opinion, traditional clothing is important.`
	grammar := `Simple present: People wear Baju Melayu during traditional ceremonies. Women wear Baju Kurung. Verb to be: It is beautiful. They are traditional clothes. Have and has: The Baju Kurung has long sleeves. The clothes have beautiful patterns. Prepositions: The selendang is around her shoulders. The samping is around his waist. Adjective plus noun: She wears a beautiful dress. They wear traditional outfit. She wears an elegant Baju Kurung.`
	tanjakMono := `Hello, friends! Take a look at this special male headdress called a Tanjak. A Tanjak is made from a piece of square songket fabric. Instead of being cut or sewn, the cloth is folded and tied into a sharp, upright crown. It shines brightly because of the gold and silver threads running through the material. Malay boys and men put this on their heads for grand events like weddings and cultural festivals. Wearing a Tanjak is not just about looking handsome — the upright tip reminds the wearer to stay brave, honest, and proud of his roots. That is why the Tanjak is so special to us in Riau. Thank you!`
	fikriSarah := `Fikri: Hey, Sarah! Look at what I am holding. Do you know what this is? Sarah: It looks like a miniature crown made of cloth. Is that a Tanjak? Fikri: Spot on! It is a traditional Malay headdress for boys and men. Sarah: The fabric looks so luxurious with all those gold threads. How do craftsmen make it so stiff and upright? Fikri: Interestingly, they do not cut or sew the cloth. They carefully fold and tie a square piece of Songket until it forms this sharp shape. Sarah: That takes real skill! When do people usually wear it? Fikri: We wear it on top of our heads for special events like weddings, cultural ceremonies, and welcoming guests of honor. Sarah: It looks very stylish. Does the pointed shape mean anything special? Fikri: Yes, absolutely! The upright tip reminds the wearer to stay brave, honorable, and true to his culture. Sarah: That is a meaningful lesson behind a handsome headpiece!`
	kainMono := `Hello, everyone. Today, I would like to talk about Kain Samping, one of the traditional clothing items in Riau Malay culture. Kain Samping is a traditional cloth worn by Malay men. It is usually worn around the waist, over a pair of trousers. It is commonly worn together with other traditional clothing, such as a baju Melayu and a tanjak. Kain Samping can have beautiful colors and patterns. The patterns may show the beauty of Malay culture. People may wear Kain Samping during traditional ceremonies, cultural events, weddings, and other special occasions. There are different ways to wear Kain Samping. The way it is folded or arranged can be part of Malay clothing traditions. Therefore, wearing Kain Samping is not only about looking beautiful. It also shows respect for Malay culture and traditions. I like Kain Samping because it is beautiful and represents the cultural identity of the Riau Malay people. As young people, we should learn about and appreciate our traditional clothing.`
	kainDialog := `Student A: Hello! What traditional clothing would you like to talk about? Student B: I would like to talk about Kain Samping. Student A: What is Kain Samping? Student B: Kain Samping is a traditional cloth worn by Malay men. Student A: Where is it worn? Student B: It is worn around the waist, usually over trousers. Student A: What traditional clothes can you wear with Kain Samping? Student B: We can wear it with baju Melayu and tanjak. Student A: When do people usually wear Kain Samping? Student B: People usually wear it during weddings, traditional ceremonies, and cultural events. Student A: What does Kain Samping look like? Student B: It has beautiful colors and patterns. Student A: Do you like Kain Samping? Student B: Yes, I do. Student A: Why do you like it? Student B: I like it because it is beautiful and represents Malay culture. Student A: Do you think young people should learn about Kain Samping? Student B: Yes, I do. Young people should learn about it because it is part of our cultural heritage.`
	ambassador := `Good morning, everyone. Today, I would like to talk about Kain Samping. Kain Samping is a traditional cloth worn by Malay men around the waist. It is important because it represents Malay culture and identity. However, one problem is that many young people rarely wear it and think it is old-fashioned. One possible reason is that modern clothes are more comfortable and follow current fashion. We can solve this problem by wearing it at cultural events and introducing it through social media. As a student, I can learn how to wear it and make a short video about it. Let's learn, appreciate, and preserve our traditional clothing. Thank you.`

	return seedChapter{
		title: "Traditional Attire",
		desc:  "Mendeskripsikan ciri fisik, bahan, corak, dan fungsi pakaian adat Melayu Riau secara lisan dalam bahasa Inggris.",
		courses: []seedCourse{
			{title: "Goal Setting", desc: "Tujuan pembelajaran yang hendak dicapai.", modules: []seedModule{
				{title: "My Learning Goals", typ: model.ModuleMonologue,
					content:    "By the end of the lesson: I can name and describe various types of Riau Malay traditional clothing and accessories in English confidently and fluently for 3 to 5 minutes. I can explain the distinct parts, wearing methods, and occasions for traditional Malay attire using correct descriptive adjectives and present tense. I can present the cultural values and philosophy of Riau Malay garments — such as modesty, dignity, and honor — clearly and accurately without reading a full script.",
					transcript: "I can name and describe various types of Riau Malay traditional clothing and accessories in English confidently and fluently for 3 to 5 minutes. I can explain the distinct parts, wearing methods, and occasions for traditional Malay attire using correct descriptive adjectives and present tense.",
					order:      1},
			}},
			{title: "Input & Exploration", desc: "Passages, vocabulary and pronunciation models.", modules: []seedModule{
				{title: "Graceful Heritage", typ: model.ModuleMonologue, content: "Read the passage about Riau Malay traditional clothing.", transcript: heritage, order: 1, tasks: []seedTask{
					speakTask("Read the passage aloud with clear pronunciation.", "heritage => heh-rih-tij; modesty => moh-des-tee; dignity => dig-nih-tee; songket => song-ket; tanjak => tahn-jahk"),
				}},
				{title: "Attire Vocabulary", typ: model.ModuleMonologue, content: "Fourteen terms from the table: Tenun songket, Tenun siak, Tekat melayu, Tanjak, Selendang, Baju cekak musang, Kopiah and songkok, Kain samping, Capal and terompah, Baju kurung teluk belanga, Tudung manto, Dokoh, Pending and sabuk, Selop and sandal. Colors: red, blue, green, yellow, black, white, brown, gold, silver, purple, pink, orange. Materials: cotton, silk, velvet, woven cloth. Adjectives: beautiful, colorful, elegant, traditional, simple, attractive, unique, comfortable, formal, modest.", transcript: attireWords, order: 2, tasks: []seedTask{
					speakTask("Name five clothing items and describe each in one sentence.", "tanjak => tahn-jahk; songkok => song-koh; selendang => seh-lehn-dahng; kebaya => keh-bah-yah; songket => song-ket; cekak musang => cheh-kahk moo-sahng"),
				}},
				{title: "Language Function", typ: model.ModuleMonologue, content: langFunc, order: 3},
				{title: "Grammar Focus", typ: model.ModuleMonologue, content: grammar, order: 4},
				{title: "Tanjak Headdress", typ: model.ModuleMonologue, content: "Listen to the monologue about the Tanjak, then read it aloud.", transcript: tanjakMono, audio: "https://drive.google.com/uc?export=download&id=1DuYnkTcAGT-enobhHinfLnjjn3ihcBBE", order: 5, tasks: []seedTask{
					speakTask("Read the Tanjak monologue aloud.", "tanjak => tahn-jahk; songket => song-ket; upright => uhp-rait; honorable => oh-noh-ruh-bul"),
				}},
				{title: "Fikri and Sarah", typ: model.ModuleDialogue, content: "Listen to the dialogue about the Tanjak and practice it with a partner.", transcript: fikriSarah, order: 6, tasks: []seedTask{
					speakTask("Choose a role and read your lines aloud.", "tanjak => tahn-jahk; luxurious => lug-zhoo-ree-us; craftsmen => krafts-men"),
				}},
			}},
			{title: "Comprehension Check", desc: "Oral questions about the Kain Samping monologue.", modules: []seedModule{
				{title: "Kain Samping Recall", typ: model.ModuleMonologue, content: "Answer each question orally in complete sentences.", transcript: kainMono, order: 1, tasks: []seedTask{
					quizTask(model.TaskOpenEnded, "What is the monologue about?", "", `{"answers":["It is about Kain Samping"]}`),
					quizTask(model.TaskOpenEnded, "Who usually wears Kain Samping?", "", `{"answers":["Malay men usually wear Kain Samping"]}`),
					quizTask(model.TaskOpenEnded, "Where is Kain Samping worn?", "", `{"answers":["It is worn around the waist"]}`),
					quizTask(model.TaskOpenEnded, "What clothing can be worn with Kain Samping?", "", `{"answers":["It can be worn with a baju Melayu and a tanjak"]}`),
					quizTask(model.TaskOpenEnded, "When do people wear Kain Samping?", "", `{"answers":["They wear it during traditional ceremonies, cultural events, weddings, and special occasions"]}`),
					quizTask(model.TaskOpenEnded, "Why is Kain Samping important in Malay culture?", "", `{"answers":["Because it represents Malay culture and identity"]}`),
					quizTask(model.TaskOpenEnded, "What can the patterns on Kain Samping show?", "", `{"answers":["They can show the beauty of Malay culture"]}`),
					quizTask(model.TaskOpenEnded, "Is Kain Samping only used for decoration? Why?", "", `{"answers":["No. It also shows respect for Malay culture and traditions"]}`),
					quizTask(model.TaskOpenEnded, "How can people wear Kain Samping?", "", `{"answers":["They can fold and arrange it in different ways"]}`),
					quizTask(model.TaskOpenEnded, "Why should young people learn about Kain Samping?", "", `{"answers":["They should learn about it to appreciate and preserve their culture"]}`),
				}},
				{title: "Personal Response", typ: model.ModuleMonologue, content: "Give your own opinion aloud. Useful expressions: I think that... In my opinion... I believe that... because... For example... I agree because... We should... Young people can...", transcript: kainMono, order: 2, tasks: []seedTask{
					// ponytail: book gives no keys for opinion answers — prompt-only OPEN_ENDED until supervisor supplies a rubric mapping.
					quizTask(model.TaskOpenEnded, "Do you like Kain Samping? Why or why not?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "When would you like to wear Kain Samping?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "What traditional clothes are popular in your area?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "How can young people help preserve traditional clothing?", "", `{"answers":[]}`),
					quizTask(model.TaskOpenEnded, "What do you think about wearing traditional clothes at cultural events?", "", `{"answers":[]}`),
					speakTask("Answer one question above for 1 minute: opinion, two reasons, one example.", "heritage => heh-rih-tij; preserve => preh-zerv"),
				}},
			}},
			{title: "Guided Speaking", desc: "Kain Samping dialogue with your own information.", modules: []seedModule{
				{title: "Kain Samping Dialogue", typ: model.ModuleDialogue, content: "Read the dialogue first. Then replace the bracketed words with your own information and practice again.", transcript: kainDialog, order: 1, tasks: []seedTask{
					speakTask("Practice the dialogue, then redo it with your own bracket answers.", "heritage => heh-rih-tij"),
					quizTask(model.TaskOpenEnded, "What traditional clothing are you talking about?", "", `{"answers":["Kain Samping"]}`),
					quizTask(model.TaskOpenEnded, "What is Kain Samping?", "", `{"answers":["A traditional cloth worn by Malay men"]}`),
					quizTask(model.TaskOpenEnded, "Where is Kain Samping worn?", "", `{"answers":["Around the waist, usually over trousers"]}`),
					quizTask(model.TaskOpenEnded, "What clothes can be worn with Kain Samping?", "", `{"answers":["Baju Melayu and tanjak"]}`),
					quizTask(model.TaskOpenEnded, "When do people usually wear it?", "", `{"answers":["Weddings, traditional ceremonies, and cultural events"]}`),
				}},
				{title: "Speaking Frame", typ: model.ModuleMonologue, content: "Close the book and speak for 3 to 4 minutes using this frame: I would like to talk about... It is a... It is worn by... It is worn around... People wear it during... It can be worn with... It has... colors and patterns. I like it because... It is important because... Young people should...", order: 2, tasks: []seedTask{
					speakTask("Speak 3 to 4 minutes using the frame without reading.", "samping => sahm-ping"),
				}},
			}},
			{title: "Independent Speaking", desc: "Opinion, problem-solution and ambassador speech.", modules: []seedModule{
				{title: "Tradition or Modern", typ: model.ModuleOralTest, content: "Many young people prefer modern clothes. Think: Why do some prefer modern clothes? Why is traditional clothing important? Should youth wear it more often? Can tradition and modern fashion combine? Choose A, B, C or All of the above and explain why: I choose... because... Compare with: Kain Samping is different from modern clothing because... Traditional clothing is... while modern clothing is... Complete the table before speaking: purpose, appearance, occasion, comfort, cultural meaning, popularity.", transcript: ambassador, order: 1, tasks: []seedTask{
					speakTask("Speak 1 to 2 minutes: your choice, two reasons, one example.", "modern => moh-dern; combine => kom-bain"),
				}},
				{title: "Cultural Day Problem", typ: model.ModuleOralTest, content: "A school plans a cultural day but students do not know how to wear traditional Malay clothing and think it is old-fashioned. Respond in 1 to 2 minutes with: Problem, possible reason, effect, solutions. The problem is... One possible reason is... This can cause... One solution is... I think the best way is...", order: 2, tasks: []seedTask{
					speakTask("Give your 1 to 2 minute problem-solution response without reading.", "solution => soh-loo-shun"),
				}},
				{title: "Young Cultural Ambassador", typ: model.ModuleOralTest, content: "Prepare a 2-minute speech: What is Kain Samping? Why is it important? What problems does it face? Why are youth less interested? How can schools and social media help? What can you personally do? Structure: Opening, description, importance, problem, reason, solution, personal action, closing. Then answer two peer questions without reading.", transcript: ambassador, order: 3, tasks: []seedTask{
					speakTask("Deliver your 2-minute ambassador speech without reading, then record.", "ambassador => am-bah-suh-der; audience => oh-dee-ens"),
				}},
			}},
		},
	}
}
