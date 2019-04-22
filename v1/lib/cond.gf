(let switch (mac (val alts..)
  (let v? (not (= val '_)))
  
  (fold (reverse alts)
        (fun (acc alt)
          (let c (head alt) p? (and v? (not (= c 'T))))
          '(if %(if p? '(%(head c) %val %(tail c)..) c)
             (do %(tail alt)..)
             %acc))
        _)))