class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        main = ""
        if (len(word1) > len(word2)):
            sub = word1[len(word2):]
            for i in range(len(word2)):
                main += word1[i]
                main += word2[i]
            main += sub
        elif (len(word2) > len(word1)):
            sub = word2[len(word1):]
            for i in range(len(word1)):
                main += word1[i]
                main += word2[i]
            main += sub
        else:
            for i in range(len(word1)):
                main += word1[i]
                main += word2[i]
        return main
